import {User} from "./user";
export class Conversation {
  id: number;
  title: string;
  user: User;
  
  constructor() {
    this.id = -1;
    this.user = new User();
  }
  
  static create(id: number, user: User): Conversation {
    let conv = new Conversation();
    conv.id = id;
    conv.user = user;
    return conv;
  }
  
}
