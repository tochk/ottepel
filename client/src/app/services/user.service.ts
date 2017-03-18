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
    //noinspection TypeScriptUnresolvedFunction
    return this.http.get('/vk/users.get?access_token=' + this.authService.token.accessToken
      + '&user_ids=' + userId)
      .map(res => {
        let body = res.json().response;
        let user = body[0];
        return User.create(user.id, user.first_name + ' ' + user.last_name, '');
      })
      .catch(err => {
        return err;
      });
  }

  getUsers(userIds:number[]):Observable<any>  {
    let reqLink = '/vk/users.get?access_token=' + this.authService.token.accessToken + '&fields=photo_100&user_ids=';
    userIds.forEach((id, i) => {
      let plus = '';
      if (i !== userIds.length - 1) {
        plus = ',';
      }
      reqLink += id + plus;
    });

    //noinspection TypeScriptUnresolvedFunction
    return this.http.get(reqLink)
      .map(res => {
        let result = [];
        let users = res.json().response;
        users.forEach(user => {
          result.push(User.create(user.uid, user.first_name + ' ' + user.last_name, user.photo_100));
        });
        return result;
      })
      .catch(err => {
        return err;
      });
  }

}
