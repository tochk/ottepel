import { Routes, RouterModule } from '@angular/router';
import {AuthComponent} from "./auth/auth.component";
import {HomeComponent} from "./home/home.component";
import {AuthGuard} from "./services/auth-guard.service";
import {PhotosComponent} from "./photos/photos.component";

const appRoutes: Routes = [
  {
    path: 'auth',
    component: AuthComponent
  },
  {
    path: 'photo/:convId',
    component: PhotosComponent
  },
  {
    path: '',
    component: HomeComponent,
    canActivate: [AuthGuard]
  },
];

export const routing = RouterModule.forRoot(appRoutes, {useHash: true});
