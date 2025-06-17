import {Component, inject, ViewEncapsulation} from '@angular/core';
import {CustomerGrpcService} from "../../../../connect/tokens";

@Component({
    selector     : 'example',
    standalone   : true,
    templateUrl  : './example.component.html',
    encapsulation: ViewEncapsulation.None,
})
export class ExampleComponent
{
    customers = inject(CustomerGrpcService)

    /**
     * Constructor
     */
    constructor(
    )
    {
        this.customers.getCustomer({id: "5c9dc1c8-1456-4bd7-af0c-7a3692ea5263"}).subscribe(
            (data)=>{
                console.error('----------------',data)
            }
        )
    }


}
