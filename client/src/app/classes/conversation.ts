import {User} from "./user";
export class Conversation {
  userId: number;
  user: User;
  charId: number;
  chatTitle: string;

  constructor() {
    this.userId = -1;
    this.user = new User();
    this.charId = -1;
    this.chatTitle = '';
  }

  static create(userId: number, user: User, charId: number, chartTitle: string): Conversation {
    let conv = new Conversation();
    conv.userId = userId;
    conv.user = user;
    conv.charId = charId;
    conv.chatTitle = chartTitle;
    return conv;
  }

}
