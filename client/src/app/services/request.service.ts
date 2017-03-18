import {Injectable} from '@angular/core';
import {Http, Response, Headers, RequestOptions} from '@angular/http';
import {Observable} from 'rxjs/Observable';
import 'rxjs/add/operator/map';
import 'rxjs/add/operator/catch';
import 'rxjs/add/observable/throw';

@Injectable()
export class RequestService {

  constructor(private http:Http) { }

  getAuthLink():Observable<any> {
    //noinspection TypeScriptUnresolvedFunction
    return this.http.get('/api/auth/')
      .map(RequestService.extractData)
      .catch(RequestService.handleError);
  }
  
  getPhotos(convId: number):Observable<any> {
    //noinspection TypeScriptUnresolvedFunction
    return this.http.get('/api/photo/' + convId)
      .map(RequestService.extractData)
      .catch(RequestService.handleError);
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
