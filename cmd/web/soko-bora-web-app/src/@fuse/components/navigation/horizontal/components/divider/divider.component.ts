import { ChangeDetectionStrategy, ChangeDetectorRef, Component, Input, OnDestroy, OnInit } from '@angular/core';
import { Subject, takeUntil } from 'rxjs';
import { FuseHorizontalNavigationComponent } from '@fuse/components/navigation/horizontal/horizontal.component';
import { FuseNavigationService } from '@fuse/components/navigation/navigation.service';
import { FuseNavigationItem } from '@fuse/components/navigation/navigation.types';

@Component({
    selector       : 'fuse-horizontal-navigation-divider-item',
    templateUrl    : './divider.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush
})
export class FuseHorizontalNavigationDividerItemComponent implements OnInit, OnDestroy
{
    // @ts-ignore
    @Input() item: FuseNavigationItem;
    @Input() name: string | null | undefined;

    private _fuseHorizontalNavigationComponent: FuseHorizontalNavigationComponent | undefined;
    private _unsubscribeAll: Subject<any> = new Subject<any>();

    /**
     * Constructor
     */
    constructor(
        private _changeDetectorRef: ChangeDetectorRef,
        private _fuseNavigationService: FuseNavigationService
    )
    {
    }

    // -----------------------------------------------------------------------------------------------------
    // @ Lifecycle hooks
    // -----------------------------------------------------------------------------------------------------

    /**
     * On init
     */
    ngOnInit(): void
    {
        // Get the parent navigation component
        this._fuseHorizontalNavigationComponent = this._fuseNavigationService.getComponent(<string>this.name);

        // Subscribe to onRefreshed on the navigation component
        this._fuseHorizontalNavigationComponent?.onRefreshed.pipe(
            takeUntil(this._unsubscribeAll)
        ).subscribe(() => {

            // Mark for check
            this._changeDetectorRef.markForCheck();
        });
    }

    /**
     * On destroy
     */
    ngOnDestroy(): void
    {
        // Unsubscribe from all subscriptions
        this._unsubscribeAll.next(null);
        this._unsubscribeAll.complete();
    }
}