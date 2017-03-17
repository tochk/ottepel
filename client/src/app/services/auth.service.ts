import {Injectable} from '@angular/core';
// import {Observable} from 'rxjs/Observable';
import {RequestService} from "./request.service";
import {Router} from '@angular/router';

@Injectable()
export class AuthService {

  // public login() {
  // }

  // public logout() {
  // }

  // public isAuth():Observable<any> {
  // }

  constructor(private requestService:RequestService,
              private router:Router) {
  }

}
