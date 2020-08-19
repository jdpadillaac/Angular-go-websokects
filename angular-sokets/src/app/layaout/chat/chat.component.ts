import {Component, OnInit} from '@angular/core';
import {ChatService} from '../../services/chat.service';

@Component({
  selector: 'app-chat',
  templateUrl: './chat.component.html',
  styleUrls: ['./chat.component.scss']
})
export class ChatComponent implements OnInit {

  public text: string;

  constructor(
    private readonly chatSrv: ChatService
  ) {
  }

  ngOnInit(): void {
    this.chatSrv.getMessages().subscribe(message => {
      console.log(message);
    });
  }

  public sendMessage(): void {
    this.chatSrv.sendMessage(this.text);
    this.text = null;
  }


}
