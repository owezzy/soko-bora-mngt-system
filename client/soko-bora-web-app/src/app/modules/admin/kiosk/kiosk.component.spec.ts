import { signal } from '@angular/core';
import { ComponentFixture, TestBed } from '@angular/core/testing';
import type { Customer } from 'proto/customerspb/api_pb';
import type { Order } from 'proto/orderingpb/api_pb';
import type {
    DemoProduct,
    DemoStore,
    GetDemoBootstrapResponse,
} from 'proto/searchpb/api_pb';
import { KioskService } from 'app/core/kiosk/kiosk.service';
import { KioskComponent } from 'app/modules/admin/kiosk/kiosk.component';

describe('KioskComponent', () => {
    let fixture: ComponentFixture<KioskComponent>;
    const orderSignal = signal<Order | null>(null);

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
            } satisfies DemoStore,
        ],
        products: [
            {
                $typeName: 'searchpb.DemoProduct',
                id: 'product-1',
                storeId: 'store-1',
                storeName: 'Fresh Harvest Grocers',
                name: 'Bananas',
            } satisfies DemoProduct,
        ],
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
        order: orderSignal.asReadonly(),
        isBusy: signal(false),
        error: signal<string | null>(null),
        statusMessage: signal('Bootstrap loaded from the backend demo configuration.'),
        storeSections: signal([
            {
                store: bootstrap.stores[0],
                products: bootstrap.products,
            },
        ]),
        loadCustomer: () => Promise.resolve(),
        addProduct: () => Promise.resolve(),
        removeProduct: () => Promise.resolve(),
        checkout: () => Promise.resolve(),
        refreshOrder: () => Promise.resolve(),
    } satisfies Partial<KioskService>;

    beforeEach(async () => {
        orderSignal.set(null);

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
        expect(text).toContain('Bananas');
        expect(text).toContain('Add a product to create a real backend basket');
    });

    it('renders the current order status when an order exists', () => {
        orderSignal.set({
            $typeName: 'orderingpb.Order',
            id: 'order-1',
            customerId: 'customer-1',
            paymentId: 'payment-1',
            items: [],
            status: 'approved',
        } satisfies Order);

        fixture.detectChanges();

        expect(fixture.nativeElement.textContent as string).toContain('Order order-1 is Approved.');
    });
});
