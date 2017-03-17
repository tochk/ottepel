export class Conversation {
  id: number;
  user: string;
  
  constructor() {
    this.id = -1;
    this.user = '';
  }
  
  static create(id: number, user: string): Conversation {
    let conv = new Conversation();
    conv.id = id;
    conv.user = user;
    return conv;
  }
  
}
