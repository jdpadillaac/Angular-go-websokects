package websockets

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
)

func SocketConnections(server *socketio.Server)  {

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("Nuevo usuario conectado con id:", s.ID())
		return nil
	})

	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("Usuario desconectado", reason)
	})
}