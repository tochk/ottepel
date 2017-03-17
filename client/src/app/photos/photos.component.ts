import { Component, OnInit } from '@angular/core';
import {Photo} from "../classes/photo";

@Component({
  selector: 'app-photos',
  templateUrl: './photos.component.html',
  styleUrls: ['./photos.component.css']
})
export class PhotosComponent implements OnInit {

  photos: Photo[];

  constructor() { }

  ngOnInit() {
  }

}
