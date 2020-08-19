import {Component, OnInit} from '@angular/core';
import {WebsocketService} from './services/websocket.service';

@Component({
  selector: 'app-root',
  templateUrl: './app.component.html',
  styleUrls: ['./app.component.scss']
})
export class AppComponent implements OnInit{

  title = 'angular-sokets';

  constructor(
    public wsService: WebsocketService
  ) {
  }

  ngOnInit(): void {}
}
