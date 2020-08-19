import {Injectable} from '@angular/core';
import {WebsocketService} from './websocket.service';
import {Observable} from 'rxjs';

@Injectable({
  providedIn: 'root'
})
export class ChatService {

  constructor(
    private wsService: WebsocketService
  ) {
  }

  public sendMessage(message: string): void {
    const payload = {
      de: 'Jonathan',
      message
    };

    this.wsService.eventEmitter('mensaje', payload);
  }


  public getMessages(): Observable<any> {
    return  this.wsService.eventListener('mensaje-nuevo');
  }
}
