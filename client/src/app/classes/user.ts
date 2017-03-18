export class User {
  id: number;
  name: string;
  
  constructor() {
    this.id = -1;
    this.name = '';
  }
  
  static create(id: number, name: string):User {
    let userObj = new User();
    userObj.id = id;
    userObj.name = name;
    return userObj;
  }
}
