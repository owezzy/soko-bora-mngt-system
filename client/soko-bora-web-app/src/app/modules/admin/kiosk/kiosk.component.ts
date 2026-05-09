import { DecimalPipe, TitleCasePipe } from '@angular/common';
import { ChangeDetectionStrategy, Component, inject, OnInit, ViewEncapsulation } from '@angular/core';
import type { Product as StoreProduct } from 'proto/storespb/api_pb';
import { KioskService } from 'app/core/kiosk/kiosk.service';

@Component({
    selector: 'kiosk',
    imports: [DecimalPipe, TitleCasePipe],
    templateUrl: './kiosk.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush,
    encapsulation: ViewEncapsulation.None,
})
export class KioskComponent implements OnInit {
    readonly kiosk = inject(KioskService);

    ngOnInit(): void {
        void this.kiosk.initialize();
    }

    addProduct(product: StoreProduct): void {
        void this.kiosk.addProduct(product);
    }

    removeProduct(productId: string): void {
        void this.kiosk.removeProduct(productId);
    }

    checkout(): void {
        void this.kiosk.checkout();
    }

    refreshOrder(): void {
        void this.kiosk.refreshOrder();
    }
}
