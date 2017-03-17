import {Component, OnInit} from '@angular/core';
import {RequestService} from "../services/request.service";
import {AuthService} from "../services/auth.service";
import {Router} from '@angular/router';
import {PATH} from "../constants/routing";

@Component({
  selector: 'app-auth',
  templateUrl: './auth.component.html',
  styleUrls: ['./auth.component.css']
})
export class AuthComponent implements OnInit {

  authLink:string;
  token:string;

  constructor(private requestService:RequestService,
              private authService:AuthService,
              private router: Router) {
  }

  ngOnInit() {
    this.token = '';
    this.getAuthLink();
  }

  getAuthLink() {
    this.requestService.getAuthLink().subscribe((res) => {
      this.authLink = res.Url;
    });
  }

  sendToken() {
    this.authService.login(this.token).then(res => {
      this.router.navigate([PATH.HOME]);
    });
  }
}
