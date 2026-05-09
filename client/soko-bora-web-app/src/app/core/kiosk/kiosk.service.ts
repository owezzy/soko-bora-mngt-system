import { ConnectError } from '@connectrpc/connect';
import { computed, inject, Injectable, signal } from '@angular/core';
import { toSignal } from '@angular/core/rxjs-interop';
import { firstValueFrom } from 'rxjs';
import { DemoBootstrapStore } from 'app/core/demo-bootstrap/demo-bootstrap.store';
import {
    BasketGrpcService,
    CustomerGrpcService,
    OrderingGrpcService,
    PaymentsGrpcService,
    StoresGrpcService,
} from 'connect/tokens';
import type { Basket, Item as BasketItem } from 'proto/basketspb/api_pb';
import type { Customer } from 'proto/customerspb/api_pb';
import type { Order as OrderingOrder } from 'proto/orderingpb/api_pb';
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
    private readonly payments = inject(PaymentsGrpcService);
    private readonly ordering = inject(OrderingGrpcService);
    private readonly stores = inject(StoresGrpcService);

    private readonly bootstrapSignal = toSignal(this.bootstrapStore.bootstrap$, {
        initialValue: this.bootstrapStore.snapshot,
    });

    private readonly customerState = signal<Customer | null>(null);
    private readonly basketState = signal<Basket | null>(null);
    private readonly paymentIdState = signal<string | null>(null);
    private readonly orderState = signal<OrderingOrder | null>(null);
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
    readonly order = computed(() => this.orderState());
    readonly paymentId = computed(() => this.paymentIdState());
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
            this.statusState.set('Demo customer and live store catalog loaded from backend services.');
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

    async checkout(): Promise<void> {
        const bootstrap = this.requireBootstrap();
        const customerId = this.requireCustomerId(bootstrap);
        const basket = this.basketState();

        if (!basket || basket.items.length === 0) {
            this.errorState.set('Add at least one product before checking out.');
            return;
        }

        const amount = this.basketTotal();

        await this.run(async () => {
            await firstValueFrom(this.customers.authorizeCustomer({ id: customerId }));

            const payment = await firstValueFrom(
                this.payments.authorizePayment({
                    customerId,
                    amount,
                }),
            );
            this.paymentIdState.set(payment.id);

            await firstValueFrom(
                this.baskets.checkoutBasket({
                    id: basket.id,
                    paymentId: payment.id,
                }),
            );

            this.statusState.set('Payment authorized and basket checked out. Waiting for order approval...');
            const order = await this.waitForOrder(basket.id);
            this.orderState.set(order);
            this.statusState.set(`Order ${order.id} is ${order.status}.`);
        });
    }

    async refreshOrder(): Promise<void> {
        const orderId = this.orderState()?.id ?? this.basketState()?.id;
        if (!orderId) {
            return;
        }

        await this.run(async () => {
            const response = await firstValueFrom(this.ordering.getOrder({ id: orderId }));
            this.orderState.set(response.order ?? null);

            if (response.order) {
                this.statusState.set(`Order ${response.order.id} is ${response.order.status}.`);
            }
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

    private async waitForOrder(orderId: string): Promise<OrderingOrder> {
        let lastError: unknown;

        for (let attempt = 0; attempt < 10; attempt += 1) {
            try {
                const response = await firstValueFrom(this.ordering.getOrder({ id: orderId }));
                if (response.order) {
                    return response.order;
                }
            } catch (error: unknown) {
                lastError = error;
            }

            await this.sleep(500);
        }

        throw lastError ?? new Error(`Order ${orderId} was not available before timeout.`);
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

    private sleep(ms: number): Promise<void> {
        return new Promise((resolve) => {
            window.setTimeout(resolve, ms);
        });
    }
}
