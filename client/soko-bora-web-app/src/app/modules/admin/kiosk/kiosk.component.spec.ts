import { signal } from '@angular/core';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import type { Customer } from 'proto/customerspb/api_pb';
import type { GetDemoBootstrapResponse, Order as SearchOrder } from 'proto/searchpb/api_pb';
import type { Product as StoreProduct, Store as MallStore } from 'proto/storespb/api_pb';
import { KioskService } from 'app/core/kiosk/kiosk.service';
import { KioskComponent } from 'app/modules/admin/kiosk/kiosk.component';

describe('KioskComponent', () => {
    let fixture: ComponentFixture<KioskComponent>;

    const bootstrap: GetDemoBootstrapResponse = {
        $typeName: 'searchpb.GetDemoBootstrapResponse',
        customer: {
            $typeName: 'searchpb.DemoCustomer',
            id: 'customer-1',
            name: 'Demo Shopper',
        },
        stores: [
            {
                $typeName: 'searchpb.DemoStore',
                id: 'store-1',
                name: 'Fresh Harvest Grocers',
            },
        ],
        products: [
            {
                $typeName: 'searchpb.DemoProduct',
                id: 'product-1',
                storeId: 'store-1',
                storeName: 'Fresh Harvest Grocers',
                name: 'Bananas',
            },
        ],
    };

    const store: MallStore = {
        $typeName: 'storespb.Store',
        id: 'store-1',
        name: 'Fresh Harvest Grocers',
        location: 'Ground Floor',
        participating: true,
    };

    const product: StoreProduct = {
        $typeName: 'storespb.Product',
        id: 'product-1',
        storeId: 'store-1',
        name: 'Bananas',
        description: 'Fresh bananas from the live catalog',
        sku: 'BNN-001',
        price: 6,
    };

    const mockKioskService = {
        bootstrap: signal(bootstrap),
        customer: signal({
            $typeName: 'customerspb.Customer',
            id: 'customer-1',
            name: 'Demo Shopper',
            smsNumber: '+254700000001',
            enabled: true,
        } satisfies Customer),
        basket: signal(null),
        basketItems: signal([]),
        basketQuantities: signal({}),
        basketTotal: signal(0),
        paymentId: signal<string | null>(null),
        order: signal<SearchOrder | null>(null),
        isBusy: signal(false),
        error: signal<string | null>(null),
        statusMessage: signal('Bootstrap loaded from the backend demo configuration.'),
        storeSections: signal([
            {
                store,
                products: [product],
            },
        ]),
        initialize: () => Promise.resolve(),
        loadCustomer: () => Promise.resolve(),
        addProduct: () => Promise.resolve(),
        removeProduct: () => Promise.resolve(),
        checkout: () => Promise.resolve(),
        refreshOrder: () => Promise.resolve(),
    } satisfies Partial<KioskService>;

    beforeEach(async () => {
        mockKioskService.basket.set(null);
        mockKioskService.basketItems.set([]);
        mockKioskService.basketQuantities.set({});
        mockKioskService.basketTotal.set(0);
        mockKioskService.paymentId.set(null);
        mockKioskService.order.set(null);
        mockKioskService.error.set(null);
        mockKioskService.statusMessage.set('Bootstrap loaded from the backend demo configuration.');

        await TestBed.configureTestingModule({
            imports: [KioskComponent],
            providers: [
                {
                    provide: KioskService,
                    useValue: mockKioskService,
                },
            ],
        }).compileComponents();

        fixture = TestBed.createComponent(KioskComponent);
        fixture.detectChanges();
    });

    it('renders the backend-backed kiosk flow', () => {
        const text = fixture.nativeElement.textContent as string;

        expect(text).toContain('Monolith kiosk demo');
        expect(text).toContain('Demo Shopper');
        expect(text).toContain('Fresh Harvest Grocers');
        expect(text).toContain('Ground Floor');
        expect(text).toContain('Bananas');
        expect(text).toContain('Fresh bananas from the live catalog');
        expect(text).toContain('SKU BNN-001');
        expect(text).toContain('Add a product to create a real backend basket');
        expect(text).toContain('Order status');
        expect(text).toContain('No order is visible yet. Check out the basket to start the saga.');
    });

    it('shows basket totals when a real basket exists', () => {
        mockKioskService.basket.set({
            $typeName: 'basketspb.Basket',
            id: 'basket-1',
            customerId: 'customer-1',
            items: [
                {
                    $typeName: 'basketspb.Item',
                    productId: 'product-1',
                    productName: 'Bananas',
                    productPrice: 6,
                    quantity: 2,
                    storeId: 'store-1',
                    storeName: 'Fresh Harvest Grocers',
                },
            ],
            checkedOut: false,
        });
        mockKioskService.basketItems.set(mockKioskService.basket().items);
        mockKioskService.basketQuantities.set({ 'product-1': 2 });
        mockKioskService.basketTotal.set(12);
        fixture.detectChanges();

        const text = fixture.nativeElement.textContent as string;

        expect(text).toContain('Basket ID');
        expect(text).toContain('basket-1');
        expect(text).toContain('2 x KES 6.00');
        expect(text).toContain('KES 12.00');
        expect(text).toContain('Authorize payment and checkout');
    });

    it('shows live order projection details after checkout', () => {
        mockKioskService.order.set({
            $typeName: 'searchpb.Order',
            orderId: 'order-1',
            customerId: 'customer-1',
            customerName: 'Demo Shopper',
            total: 12,
            status: 'Ready For Pickup',
            items: [
                {
                    $typeName: 'searchpb.Order.Item',
                    productId: 'product-1',
                    storeId: 'store-1',
                    productName: 'Bananas',
                    storeName: 'Fresh Harvest Grocers',
                    price: 6,
                    quantity: BigInt(2),
                },
            ],
        });

        fixture.detectChanges();

        const text = fixture.nativeElement.textContent as string;

        expect(text).toContain('Order order-1 is Ready For Pickup.');
        expect(text).toContain('Projected basket total');
        expect(text).toContain('KES 12.00');
        expect(text).toContain('Ready For Pickup');
    });
});
