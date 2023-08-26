import { NgModule } from '@angular/core';
import { RouterModule } from '@angular/router';
import { MatButtonModule } from '@angular/material/button';
import { MatCheckboxModule } from '@angular/material/checkbox';
import { MatFormFieldModule } from '@angular/material/form-field';
import { MatIconModule } from '@angular/material/icon';
import { MatInputModule } from '@angular/material/input';
import { MatProgressSpinnerModule } from '@angular/material/progress-spinner';
import { FuseCardModule } from '../../../../@fuse/components/card';
import { FuseAlertModule } from '../../../../@fuse/components/alert';
import { SharedModule } from '../../../shared/shared.module';
import { Auth0Component } from './auth0.component';
import { auth0Routes } from './auth0.routing';
import { AuthButtonComponent } from './auth0.button.component';
import { FuseLoadingBarModule } from '../../../../@fuse/components/loading-bar';

@NgModule({
  declarations: [Auth0Component, AuthButtonComponent],
  imports: [
    RouterModule.forChild(auth0Routes),
    MatButtonModule,
    MatCheckboxModule,
    MatFormFieldModule,
    MatIconModule,
    MatInputModule,
    MatProgressSpinnerModule,
    FuseCardModule,
    FuseAlertModule,
    SharedModule,
    FuseLoadingBarModule,
  ],
})
export class Auth0Module {}
