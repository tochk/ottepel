export class Photo {
  id: number;
  url: string;
  
  constructor() {
    
  }
  
  static create(id: number, url: string): Photo {
    let photoObj = new Photo();
    photoObj.id = id;
    photoObj.url = url;
    return photoObj;
  }

}
