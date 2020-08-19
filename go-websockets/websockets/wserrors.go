package websockets

import (
	"fmt"
	socketio "github.com/googollee/go-socket.io"
)

func ErrorConnections (server *socketio.Server) {
	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("meet error:", e)
	})
}