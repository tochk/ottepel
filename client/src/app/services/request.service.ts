import {Injectable} from '@angular/core';
import {Http, Response, Headers, RequestOptions} from '@angular/http';
import {Observable} from 'rxjs/Observable';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';
import 'rxjs/add/observable/interval';
import 'rxjs/Rx';
import {AuthService} from "./auth.service";
import {Photo} from "../classes/photo";

@Injectable()
export class RequestService {

  constructor(private http:Http, private authService:AuthService) {
  }

  getAuthLink():Observable<any> {
    //noinspection TypeScriptUnresolvedFunction
    return this.http.get('/api/auth/')
      .map(RequestService.extractData)
      .catch(RequestService.handleError);
  }

  getPhotos(convId:number):Observable<any> {
    let body = JSON.stringify({
      "AccessToken": this.authService.token.accessToken,
      "UserId": convId
    });

    let headers = new Headers({'Content-Type': 'application/json'});
    let options = new RequestOptions({headers: headers});

    //noinspection TypeScriptUnresolvedFunction
    return this.http.post('/api/getPhotos/', body, options)
      .map((res:Response) => {
        let result:Photo[] = [];
        let body = res.json();
        body.Photos.forEach(ph => {
          result.push(new Photo(ph));
        });
        return result;
      })
      .catch(RequestService.handleError);
  }

  getTokenForArchive(photos:string[]):Observable<any> {
    let body = JSON.stringify({
      "AccessToken": this.authService.token.accessToken,
      "Photos": photos
    });

    let headers = new Headers({'Content-Type': 'application/json'});
    let options = new RequestOptions({headers: headers});

    //noinspection TypeScriptUnresolvedFunction
    return this.http.post('/api/getArchive/', body, options)
      .map((res:Response) => {
        let body = res.json();
        return body.Token;
      })
      .catch(RequestService.handleError);
  }

  loadArchive(token:string):Observable<any> {
    //noinspection TypeScriptUnresolvedFunction
    return this.http.get('api/userFiles/' + token + '.zip')
      .map(RequestService.extractData)
      .catch(RequestService.handleError);
  }

  isFileExist(token:string):Observable<any> {
    let body = JSON.stringify({
      "Token": token,
    });

    let headers = new Headers({'Content-Type': 'application/json'});
    let options = new RequestOptions({headers: headers});

    return Observable.interval(1000)
      .switchMap(() => this.http.post('/api/isFileExist/', body, options))
      .map(res => {
        let body = res.json();
        return body.IsExist;
      });
  }

  private static extractData(res:Response) {
    let body = res.json();
    return body || {};
  }

  private static handleError(error:any) {
    let errMsg = (error.message) ? error.message :
      error.status ? `${error.status} - ${error.statusText}` : 'Server error';
    console.error('error: ' + errMsg);
    return Observable.throw(errMsg);
  }
}
