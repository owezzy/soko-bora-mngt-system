import { ConnectError } from '@connectrpc/connect';
import { computed, inject, Injectable, signal } from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';
import { firstValueFrom } from 'rxjs';
import { DemoBootstrapStore } from 'app/core/demo-bootstrap/demo-bootstrap.store';
import { BasketGrpcService, CustomerGrpcService, StoresGrpcService } from 'connect/tokens';
import type { Basket, Item as BasketItem } from 'proto/basketspb/api_pb';
import type { Customer } from 'proto/customerspb/api_pb';
import type { GetDemoBootstrapResponse } from 'proto/searchpb/api_pb';
import type { Product as StoreProduct, Store as MallStore } from 'proto/storespb/api_pb';

interface KioskStoreSection {
    store: MallStore;
    products: StoreProduct[];
}

@Injectable({ providedIn: 'root' })
export class KioskService {
    private readonly bootstrapStore = inject(DemoBootstrapStore);
    private readonly customers = inject(CustomerGrpcService);
    private readonly baskets = inject(BasketGrpcService);
    private readonly stores = inject(StoresGrpcService);

    private readonly bootstrapSignal = toSignal(this.bootstrapStore.bootstrap$, {
        initialValue: this.bootstrapStore.snapshot,
    });

    private readonly customerState = signal<Customer | null>(null);
    private readonly basketState = signal<Basket | null>(null);
    private readonly busyState = signal(false);
    private readonly errorState = signal<string | null>(null);
    private readonly statusState = signal('Bootstrap loaded from the backend demo configuration.');
    private readonly storeSectionsState = signal<KioskStoreSection[]>([]);

    readonly bootstrap = computed<GetDemoBootstrapResponse | null>(() => this.bootstrapSignal());
    readonly customer = computed(() => this.customerState());
    readonly basket = computed(() => this.basketState());
    readonly basketItems = computed(() => this.basketState()?.items ?? []);
    readonly basketQuantities = computed<Record<string, number>>(() => {
        const quantities: Record<string, number> = {};

        for (const item of this.basketItems()) {
            quantities[item.productId] = item.quantity;
        }

        return quantities;
    });
    readonly basketTotal = computed(() => this.calculateBasketTotal(this.basketItems()));
    readonly isBusy = computed(() => this.busyState());
    readonly error = computed(() => this.errorState());
    readonly statusMessage = computed(() => this.statusState());
    readonly storeSections = computed<KioskStoreSection[]>(() => this.storeSectionsState());

    async initialize(): Promise<void> {
        await this.run(async () => {
            const bootstrap = this.requireBootstrap();
            const customerId = this.requireCustomerId(bootstrap);

            const [customerResponse, storesResponse] = await Promise.all([
                firstValueFrom(this.customers.getCustomer({ id: customerId })),
                firstValueFrom(this.stores.getParticipatingStores({})),
            ]);

            this.customerState.set(customerResponse.customer ?? null);

            const stores = storesResponse.stores;
            const catalogs = await Promise.all(
                stores.map(async (store) => {
                    const catalogResponse = await firstValueFrom(this.stores.getCatalog({ storeId: store.id }));

                    return {
                        store,
                        products: catalogResponse.products,
                    } satisfies KioskStoreSection;
                }),
            );

            this.storeSectionsState.set(catalogs);
            this.statusState.set('Demo customer, live store catalog, and basket-ready kiosk flow loaded from backend services.');
        });
    }

    async loadCustomer(): Promise<void> {
        if (this.customerState()) {
            return;
        }

        await this.initialize();
    }

    async addProduct(product: StoreProduct): Promise<void> {
        const bootstrap = this.requireBootstrap();
        const customerId = this.requireCustomerId(bootstrap);

        await this.run(async () => {
            const basketId = await this.ensureBasket(customerId);

            await firstValueFrom(
                this.baskets.addItem({
                    id: basketId,
                    productId: product.id,
                    quantity: 1,
                }),
            );

            await this.refreshBasket();
            this.statusState.set(`Added ${product.name} through the Basket service.`);
        });
    }

    async removeProduct(productId: string): Promise<void> {
        const basket = this.basketState();
        if (!basket) {
            return;
        }

        const existingItem = basket.items.find((item) => item.productId === productId);
        if (!existingItem) {
            return;
        }

        await this.run(async () => {
            await firstValueFrom(
                this.baskets.removeItem({
                    id: basket.id,
                    productId,
                    quantity: 1,
                }),
            );

            await this.refreshBasket();
            this.statusState.set(`Updated basket contents through the Basket service.`);
        });
    }

    private async ensureBasket(customerId: string): Promise<string> {
        const existingBasket = this.basketState();
        if (existingBasket) {
            return existingBasket.id;
        }

        const response = await firstValueFrom(this.baskets.startBasket({ customerId }));
        await this.refreshBasket(response.id);
        return response.id;
    }

    private async refreshBasket(basketId = this.basketState()?.id): Promise<void> {
        if (!basketId) {
            return;
        }

        const response = await firstValueFrom(this.baskets.getBasket({ id: basketId }));
        this.basketState.set(response.basket ?? null);
    }

    private requireBootstrap(): GetDemoBootstrapResponse {
        const bootstrap = this.bootstrap();
        if (!bootstrap) {
            throw new Error('The kiosk bootstrap configuration has not loaded yet.');
        }

        return bootstrap;
    }

    private requireCustomerId(bootstrap: GetDemoBootstrapResponse): string {
        const customerId = bootstrap.customer?.id;
        if (!customerId) {
            throw new Error('The backend bootstrap configuration did not include a demo customer.');
        }

        return customerId;
    }

    private calculateBasketTotal(items: BasketItem[]): number {
        return items.reduce((total, item) => total + item.productPrice * item.quantity, 0);
    }

    private async run(operation: () => Promise<void>): Promise<void> {
        this.busyState.set(true);
        this.errorState.set(null);

        try {
            await operation();
        } catch (error: unknown) {
            this.errorState.set(this.describeError(error));
        } finally {
            this.busyState.set(false);
        }
    }

    private describeError(error: unknown): string {
        if (error instanceof ConnectError) {
            return error.rawMessage;
        }

        if (error instanceof Error) {
            return error.message;
        }

        return 'The kiosk request failed unexpectedly.';
    }
}
