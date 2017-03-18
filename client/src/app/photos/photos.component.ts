import { Component, OnInit } from '@angular/core';
import {Photo} from "../classes/photo";
import {ActivatedRoute, Params} from '@angular/router';
import {RequestService} from "../services/request.service";

@Component({
  selector: 'app-photos',
  templateUrl: './photos.component.html',
  styleUrls: ['./photos.component.css']
})
export class PhotosComponent implements OnInit {

  photos: Photo[];

  constructor(private route:ActivatedRoute,
              private requestService: RequestService) { }

  ngOnInit() {
    this.route.params.forEach((params:Params) => {
      let convId = +params['convId'];
      this.requestService.getPhotos(convId).subscribe((res) => {
        this.photos = res;
      });
    });

    // this.photos = [];
    // for (let i = 0; i < 10; i++)
    //   this.photos.push(new Photo("https://pp.userapi.com/c639825/v639825919/f168/LRPf8ZzAGCk.jpg"));
  }

}
