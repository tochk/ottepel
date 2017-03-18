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

  countConv:number;
  conv:Conversation[];

  constructor(private http:Http, private authService:AuthService,
              private userService:UserService) {
  }

  getConversation(offset:number, step: number) {
    //noinspection TypeScriptUnresolvedFunction
    return new Promise((resolve, reject) => {
      this.getConv(offset, step).subscribe(res => {
        this.conv = res.chats;

        if (res.userIds.length > 0) {
          this.userService.getUsers(res.userIds).subscribe(users => {
            users.forEach(user => {
              let localConv = new Conversation();
              localConv.user = user;
              localConv.userId = user.id;
              this.conv.push(localConv);
            });
            resolve(this.conv);
          });
        } else {
          resolve([]);
        }

      });
    });
  }

  getConv(offset:number, step: number):Observable<any> {
    if (offset === 0) {
      this.conv = [];
    }
    //noinspection TypeScriptUnresolvedFunction
    return this.http.get('/vk/messages.getDialogs?access_token=' + this.authService.token.accessToken + '&offset=' + offset
      + '&count=' + step + '&v=5.62')
      .map(res => {
        let body = res.json().response;
        this.countConv = body.cout;

        let chats:Conversation[] = [];
        let userIds = [];
        body.items.forEach(item => {
          let mess = item.message;
          if (mess.chat_id) {
            let localConv = new Conversation();
            localConv.chatId = mess.chat_id + 2000000000;
            localConv.chatTitle = mess.title;
            chats.push(localConv);
          } else {
            userIds.push(mess.user_id);
          }
        });

        return {'chats': chats, 'userIds': userIds};
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
