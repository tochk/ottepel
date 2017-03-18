import {Injectable} from '@angular/core';
import {Http, Response, Headers, RequestOptions} from '@angular/http';
import {Observable} from 'rxjs/Observable';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';
import {Token} from "../classes/token";
import {User} from "../classes/user";
import {LOCAL_STORAGE} from "../constants/local-storage";

@Injectable()
export class AuthService {

  isAuthorize:boolean;
  user:User;
  token:Token;

  constructor(private http:Http) {
  }

  public login(token:string) {
    return new Promise((resolve, reject) => {
      this.setToken(token).subscribe(res => {
        this.getUserInfo().subscribe(res => {
          resolve(this.isAuthorize = true);
        });
      });
    });
  }

  public logout() {
    this.isAuthorize = false;
    localStorage.removeItem(LOCAL_STORAGE.ACCESS_TOKEN);
    localStorage.removeItem(LOCAL_STORAGE.EXPIRES_IN);
  }

  public isAuth() {
    return new Promise((resolve, reject) => {
      let token = localStorage.getItem(LOCAL_STORAGE.ACCESS_TOKEN),
        time = +localStorage.getItem(LOCAL_STORAGE.EXPIRES_IN);
      if (token && time >= 0) {
        this.token = new Token(token, +time);
        this.getUserInfo().subscribe(
          res => {
            if (res) {
              resolve(this.isAuthorize = true);
            }
          },
          error => {
            resolve(this.isAuthorize = false);
          });
      } else {
        resolve(this.isAuthorize = false);
      }
    });
  }

  setToken(token:string):Observable<any> {
    let body = JSON.stringify({
      "Url": token
    });

    let headers = new Headers({'Content-Type': 'application/json'});
    let options = new RequestOptions({headers: headers});

    //noinspection TypeScriptUnresolvedFunction
    return this.http.post('/api/token/', body, options)
      .map((res:Response) => {
        let body = res.json() || {};

        this.token = new Token(body.AccessToken, +body.ExpiresIn);

        localStorage.setItem(LOCAL_STORAGE.ACCESS_TOKEN, this.token.accessToken);
        localStorage.setItem(LOCAL_STORAGE.EXPIRES_IN, this.token.expiresIn.toString());

        return true;
      })
      .catch(AuthService.handleError);
  }

  getUserInfo():Observable<any> {
    //noinspection TypeScriptUnresolvedFunction
    return this.http.get('/vk/account.getProfileInfo?access_token=' + this.token.accessToken)
      .map((res:Response) => {
        let resBody:any = res.json().response;
        this.user = new User();
        this.user.name = resBody.first_name + " " + resBody.last_name;
        return true;
      })
      .catch(AuthService.handleError);
  }

  private static handleError(error:any) {
    let errMsg = (error.message) ? error.message :
      error.status ? `${error.status} - ${error.statusText}` : 'Server error';
    console.error('error: ' + errMsg);
    return Observable.throw(errMsg);
  }

}
