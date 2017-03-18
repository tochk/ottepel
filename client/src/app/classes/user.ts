export class User {
  id: number;
  name: string;
  avatar: string;
  
  constructor() {
    this.id = -1;
    this.name = '';
    this.avatar = '';
  }
  
  static create(id: number, name: string, ava: string):User {
    let userObj = new User();
    userObj.id = id;
    userObj.name = name;
    userObj.avatar = ava;
    return userObj;
  }
}
