import {Injectable} from '@angular/core';
import {Http, Response, Headers, RequestOptions} from '@angular/http';
import {Observable} from 'rxjs/Observable';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';
import {Conversation} from "../classes/conversation";
import {AuthService} from "./auth.service";
import {UserService} from "./user.service";
import {User} from "../classes/user";

@Injectable()
export class ConversationService {

  // countConv:number;
  conv:Conversation[];

  constructor(private http:Http, private authService:AuthService,
              private userService:UserService) {
  }

  getConv(limit: number) {
    this.conv = [];
    return this.http.get('/vk/messages.searchDialogs?access_token=' + this.authService.token.accessToken + '&limit=' + limit + '&v=5.62')
      .map(res => {
        let body = res.json().response;
        body.forEach(item => {
          let c:Conversation = new Conversation();
          c.id = +item.id;
          if (item.type === "chat") {
            c.id += 2000000000;
            c.title = item.title;
          } else {
            c.user = User.create(-1, item.first_name + ' ' + item.last_name);
          }

          this.conv.push(c);
        });
        return this.conv;
      })
      .catch(ConversationService.handleError);
  }

  private static handleError(error:any) {
    let errMsg = (error.message) ? error.message :
      error.status ? `${error.status} - ${error.statusText}` : 'Server error';
    console.error('error: ' + errMsg);
    return Observable.throw(errMsg);
  }
}
