export class User {
  name: string;
  
  constructor() {
    this.name = '';
  }
  
  static create(name: string): User {
    let userObj = new User();
    userObj.name = name;
    return userObj;
  }
}
