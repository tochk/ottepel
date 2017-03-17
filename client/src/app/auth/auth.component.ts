import { Component, OnInit } from '@angular/core';
import {RequestService} from "../services/request.service";

@Component({
  selector: 'app-auth',
  templateUrl: './auth.component.html',
  styleUrls: ['./auth.component.css']
})
export class AuthComponent implements OnInit {

  authLink: string;
  token: string;

  constructor(private requestService: RequestService) { }

  ngOnInit() {
    this.requestService.getAuthLink().subscribe((res) => {
      this.authLink = res.Url;
    });
    this.token = '';
  }

  getToken() {
  }

  sendToken() {
    this.requestService.setToken(this.token).subscribe((res) => {
    });
  }

}
