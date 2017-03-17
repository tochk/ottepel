import { Routes, RouterModule } from '@angular/router';
import {AuthComponent} from "./auth/auth.component";
import {HomeComponent} from "./home/home.component";
import {AuthGuard} from "./services/auth-guard.service";

const appRoutes: Routes = [
  {
    path: 'auth',
    component: AuthComponent
  },
  {
    path: '',
    component: HomeComponent,
    canActivate: [AuthGuard]
  },
];

export const routing = RouterModule.forRoot(appRoutes, {useHash: true});
