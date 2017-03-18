import { Component, OnInit } from '@angular/core';
import {Conversation} from "../classes/conversation";
import {ConversationService} from "../services/conversation.service";

@Component({
  selector: 'app-conversation',
  templateUrl: './conversation.component.html',
  styleUrls: ['./conversation.component.css']
})
export class ConversationComponent implements OnInit {

  conversations: Conversation[];

  constructor(private convService: ConversationService) { }

  ngOnInit() {
    this.convService.getConversation(0).then((res: any) => {
      this.conversations = res;
    });
  }

}
