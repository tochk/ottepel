import {Component, OnInit, HostListener} from '@angular/core';
import {Conversation} from "../classes/conversation";
import {ConversationService} from "../services/conversation.service";
import {PATH} from "../constants/routing";

@Component({
  selector: 'app-conversations',
  templateUrl: './conversations.component.html',
  styleUrls: ['./conversations.component.css']
})
export class ConversationsComponent implements OnInit {

  isEnd:boolean;
  offset:number;
  step:number;
  conversations:Conversation[];

  constructor(private convService:ConversationService) {
  }

  ngOnInit() {
    this.isEnd = false;
    this.offset = 0;
    this.step = 20;
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

  getTooltipData(convIndex: number) {
    return this.conversations[convIndex].userId === -1 ?
      this.conversations[convIndex].chatTitle : this.conversations[convIndex].user.name;
  }

  getRouterLink(convIndex: number) {
    let res = [PATH.PHOTO];
    let id = this.conversations[convIndex].userId === -1 ?
      this.conversations[convIndex].chatId : this.conversations[convIndex].userId;
    res.push(id.toString());
    return res;
  }

}
