import { Component, OnInit } from '@angular/core';
import { AuthService } from '@auth0/auth0-angular';

@Component({
  selector: 'soko-bora-web-app-auth0',
  template: '<auth0-button></auth0-button>',
  styles: [],
})
export class Auth0Component implements OnInit {
  constructor(private auth0: AuthService) {
    // this.auth0.loginWithRedirect();
  }

  ngOnInit(): void {
    return;
  }
}
