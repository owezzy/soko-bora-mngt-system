import { Injectable } from '@angular/core';
import type { GetDemoBootstrapResponse } from '../../../proto/searchpb/api_pb';
import { BehaviorSubject, Observable } from 'rxjs';

@Injectable({ providedIn: 'root' })
export class DemoBootstrapStore {
    private readonly _bootstrap = new BehaviorSubject<GetDemoBootstrapResponse | null>(null);

    get bootstrap$(): Observable<GetDemoBootstrapResponse | null> {
        return this._bootstrap.asObservable();
    }

    get snapshot(): GetDemoBootstrapResponse | null {
        return this._bootstrap.getValue();
    }

    setBootstrap(bootstrap: GetDemoBootstrapResponse): void {
        this._bootstrap.next(bootstrap);
    }
}
