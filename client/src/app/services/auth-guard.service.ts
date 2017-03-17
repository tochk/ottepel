import { Injectable }     from '@angular/core';
import { CanActivate, Router }    from '@angular/router';
import { AuthService } from './auth.service';
import {PATH} from "../constants/routing";

@Injectable()
export class AuthGuard implements CanActivate {

  constructor(private authService: AuthService,
              private router: Router) {
  }

  canActivate() {
    if (this.authService.isAuthorize) {
      return true;
    }
    
    this.router.navigate([PATH.AUTH]);
    return false;
  }
}
