import {Component, OnInit, HostListener} from '@angular/core';
import {Photo} from "../classes/photo";
import {ActivatedRoute, Params} from '@angular/router';
import {RequestService} from "../services/request.service";
import {URLS} from "../constants/urls";

@Component({
  selector: 'app-photos',
  templateUrl: './photos.component.html',
  styleUrls: ['./photos.component.css']
})
export class PhotosComponent implements OnInit {
  photos:Photo[];
  allPhoto:Photo[];
  selectedPhoto:boolean[];
  countSelectPhoto:number;
  lastIndex:number;
  step:number;

  constructor(private route:ActivatedRoute,
              private requestService:RequestService) {
  }

  ngOnInit() {
    this.step = 20;
    this.countSelectPhoto = 0;
    this.route.params.forEach((params:Params) => {
      let convId = +params['convId'];
      this.requestService.getPhotos(convId).subscribe((res) => {
        let len = res.length;
        this.allPhoto = res;

        this.selectedPhoto = new Array(len);
        //noinspection TypeScriptUnresolvedFunction
        this.selectedPhoto.fill(false);

        this.lastIndex = len < this.step ? len : this.step;
        this.photos = this.allPhoto.slice(0, this.lastIndex);
      });
    });
  }

  addPhoto() {
    if (this.allPhoto) {
      let li = this.allPhoto.length < this.lastIndex + this.step ? this.allPhoto.length : this.lastIndex + this.step;
      for (let i = this.lastIndex; i < li; i++) {
        this.photos.push(this.allPhoto[i]);
      }
      this.lastIndex = li;
    }
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
      this.addPhoto();
    }
  }

  selectPhoto(photoIndex:number) {
    this.selectedPhoto[photoIndex] = !this.selectedPhoto[photoIndex];
    if (this.selectedPhoto[photoIndex]) {
      this.countSelectPhoto++;
    } else {
      this.countSelectPhoto--;
    }
  }

  getArchive() {
    let photoForArchive = [];
    this.selectedPhoto.forEach((photo, i) => {
      if (photo) {
        photoForArchive.push(this.allPhoto[i].url);
      }
    });

    this.requestService.getTokenForArchive(photoForArchive).subscribe(token => {
      window.open(URLS.SERVER + '/userFiles/' + token + '.zip');
    });
  }
}
