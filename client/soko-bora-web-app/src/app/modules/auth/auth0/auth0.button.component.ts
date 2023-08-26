import { Component, Inject } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';
import { DOCUMENT } from '@angular/common';

@Component({
  selector: 'auth0-button',
  template: `
    <ng-container *ngIf="auth.isAuthenticated$ | async; else loggedOut">
      <button
        mat-raised-button
        color="warn"
        (click)="
          auth.logout({
            logoutParams: {
              returnTo: window.location.origin + '/auth0-sign-in'
            }
          })
        "
      >
        Log out
      </button>
    </ng-container>

    <ng-template #loggedOut>
      <span *ngIf="auth.isLoading$ | async as loading; else loaded">
        Loading
      </span>
      <ng-template #loaded>
        <button
          color="primary"
          mat-raised-button
          (click)="auth.loginWithRedirect()"
        >
          Log in
        </button>
      </ng-template>
    </ng-template>
  `,
  styles: [],
})
export class AuthButtonComponent {
  constructor(
    @Inject(DOCUMENT) public document: Document,
    public auth: AuthService
  ) {}

  protected readonly window = window;
}
