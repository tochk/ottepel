import {Component, OnInit} from '@angular/core';
import {RequestService} from "../services/request.service";
import {Conversation} from "../classes/conversation";

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  conversations: Conversation[];

  constructor(private requestService: RequestService) { }

  ngOnInit() {
    // this.requestService.getChats().subscribe(res => {
    // });

    this.conversations = [];
    for (let i = 0; i < 10; i++) {
      this.conversations.push(Conversation.create(i, 'Some user name'));
    }
  }
}
