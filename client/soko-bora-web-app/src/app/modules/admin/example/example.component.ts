import { AsyncPipe } from '@angular/common';
import { ChangeDetectionStrategy, Component, inject, ViewEncapsulation } from '@angular/core';
import { DemoBootstrapStore } from 'app/core/demo-bootstrap/demo-bootstrap.store';

@Component({
    selector     : 'example',
    imports      : [AsyncPipe],
    templateUrl  : './example.component.html',
    changeDetection: ChangeDetectionStrategy.OnPush,
    encapsulation: ViewEncapsulation.None,
})
export class ExampleComponent {
    protected readonly bootstrap$ = inject(DemoBootstrapStore).bootstrap$;
}
