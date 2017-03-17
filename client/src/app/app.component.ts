import { Component, OnInit } from '@angular/core';
import {AuthService} from "./services/auth.service";
import {Router} from '@angular/router';
import {PATH} from "./constants/routing";

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.css']
})
export class AppComponent implements OnInit {

  constructor(private authService: AuthService,
              private router: Router) {}

  ngOnInit() {
    let nav = window.location.hash.replace('#', '');

    if (nav === PATH.AUTH) {
      nav = '/';
    }

    this.authService.isAuth().then(isAuth => {
      if (isAuth) {
        this.router.navigate([nav]);
      } else {
        this.router.navigate([PATH.AUTH]);
      }
    })
  }
}
