package websockets

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"github.com/mitchellh/mapstructure"
)

type person struct {
	De string
	Message string
}

func SocketEvents(server *socketio.Server)  {

	server.OnEvent("/", "mensaje", func(s socketio.Conn, msg interface{}) {

		var result person
		_ = mapstructure.Decode(msg, &result)
		fmt.Println(result)
		s.Emit("mensaje-nuevo", "Holaaa")

	})

	server.OnEvent("/chat", "msg", func(s socketio.Conn, msg string) string {
		s.SetContext(msg)
		return "recv " + msg
	})
}