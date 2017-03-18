import {Component, OnInit, HostListener} from '@angular/core';
import {Conversation} from "../classes/conversation";
import {ConversationService} from "../services/conversation.service";

@Component({
  selector: 'app-conversation',
  templateUrl: './conversation.component.html',
  styleUrls: ['./conversation.component.css']
})
export class ConversationComponent implements OnInit {

  isEnd:boolean;
  offset:number;
  step:number;
  conversations:Conversation[];

  constructor(private convService:ConversationService) {
  }

  ngOnInit() {
    this.isEnd = false;
    this.offset = 0;
    this.step = 200;
    this.conversations = [];
    this.loadConv();
  }

  @HostListener("window:scroll", ['$event'])
  onWindowScroll(event) {
    let windowHeight = "innerHeight" in window ? window.innerHeight
      : document.documentElement.offsetHeight;
    let body = document.body, html = document.documentElement;
    let docHeight = Math.max(body.scrollHeight,
      body.offsetHeight, html.clientHeight,
      html.scrollHeight, html.offsetHeight);
    let windowBottom = windowHeight + window.pageYOffset;
    if (windowBottom >= docHeight) {
      this.loadConv();
    }
  }

  loadConv() {
    if (!this.isEnd) {
      this.convService.getConversation(this.offset, this.step).then((res:any) => {
        res.forEach(res => {
          this.conversations.push(res);
        });
        if (res.length === 0) {
          this.isEnd = true;
        }
        this.offset += this.step;
      });
    }
  }

}
