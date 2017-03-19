import {User} from "./user";
export class Conversation {
  userId: number;
  user: User;
  chatId: number;
  chatTitle: string;
  chatImg: string;

  constructor() {
    this.userId = -1;
    this.user = new User();
    this.chatId = -1;
    this.chatTitle = '';
    this.chatImg = 'https://vk.com/images/camera_100.png';
  }

  static create(userId: number, user: User, chatId: number, chartTitle: string): Conversation {
    let conv = new Conversation();
    conv.userId = userId;
    conv.user = user;
    conv.chatId = chatId;
    conv.chatTitle = chartTitle;
    return conv;
  }

}
