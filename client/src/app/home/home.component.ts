import {Component, OnInit} from '@angular/core';
import {Conversation} from "../classes/conversation";
import {ConversationService} from "../services/conversation.service";

@Component({
  selector: 'app-home',
  templateUrl: './home.component.html',
  styleUrls: ['./home.component.css']
})
export class HomeComponent implements OnInit {

  conversations: Conversation[];

  constructor(private convService: ConversationService) { }

  ngOnInit() {
    this.convService.getConv(200).subscribe((res: any) => {
      this.conversations = res;
    });
  }
}
