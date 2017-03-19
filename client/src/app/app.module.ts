import { BrowserModule } from '@angular/platform-browser';
import { NgModule } from '@angular/core';
import { FormsModule } from '@angular/forms';
import { HttpModule, JsonpModule } from '@angular/http';

import { AppComponent } from './app.component';
import { AuthComponent } from './auth/auth.component';
import { HomeComponent } from './home/home.component';

import { routing } from "./app.routing";
import {RequestService} from "./services/request.service";
import {AuthGuard} from "./services/auth-guard.service";
import {AuthService} from "./services/auth.service";
import { PhotosComponent } from './photos/photos.component';
import {ConversationService} from "./services/conversation.service";
import {UserService} from "./services/user.service";
import { ConversationsComponent } from './conversations/conversations.component';
import { MaterializeModule } from 'angular2-materialize';
import { SidenavComponent } from './sidenav/sidenav.component';
import { LoadingCircleComponent } from './loading-circle/loading-circle.component';

@NgModule({
  declarations: [
    AppComponent,
    AuthComponent,
    HomeComponent,
    PhotosComponent,
    ConversationsComponent,
    SidenavComponent,
    LoadingCircleComponent,
  ],
  imports: [
    BrowserModule,
    FormsModule,
    HttpModule,
    JsonpModule,
    MaterializeModule,
    routing,
  ],
  providers: [
    RequestService,
    AuthGuard,
    AuthService,
    ConversationService,
    UserService
  ],
  bootstrap: [AppComponent]
})
export class AppModule { }
