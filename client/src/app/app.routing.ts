import { Routes, RouterModule } from '@angular/router';
import {AuthComponent} from "./auth/auth.component";
import {HomeComponent} from "./home/home.component";

const appRoutes: Routes = [
  {
    path: 'auth',
    component: AuthComponent
  },
  {
    path: '',
    component: HomeComponent
  }
];

export const routing = RouterModule.forRoot(appRoutes, {useHash: true});
