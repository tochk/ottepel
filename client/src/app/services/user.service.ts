import {Injectable} from '@angular/core';
import {Http, Response, Headers, RequestOptions} from '@angular/http';
import {Observable} from 'rxjs/Observable';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';
import {AuthService} from "./auth.service";
import {User} from "../classes/user";

@Injectable()
export class UserService {

  constructor(private http:Http, private authService:AuthService) {
  }

  getUser(userId:number):Observable<any> {
    return this.http.get('/vk/users.get?access_token=' + this.authService.token.accessToken
      + '&user_ids=' + userId)
      .map(res => {
        let body = res.json().response;
        let user = body[0];
        return User.create(user.id, user.first_name + ' ' + user.last_name);
      })
      .catch(err => {
        return err;
      });
  }

}
