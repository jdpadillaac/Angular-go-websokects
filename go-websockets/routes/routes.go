package routes

import (
	"fmt"
	engineio "github.com/googollee/go-engine.io"
	"github.com/googollee/go-engine.io/transport"
	"github.com/googollee/go-engine.io/transport/polling"
	"github.com/googollee/go-engine.io/transport/websocket"
	socketio "github.com/googollee/go-socket.io"
	"github.com/jdpadillaac/Angular-go-websokects/tree/master/go-websockets/websockets"
	"github.com/rs/cors"
	"log"
	"net/http"
)

func AppRoutesHandler() {

	fmt.Println("Implementacionde servidor")

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte("{\"hello\": \"world\"}"))
	})

	pt := polling.Default
	wt := websocket.Default
	wt.CheckOrigin = func(req *http.Request) bool {
		return true
	}

	server, err := socketio.NewServer(&engineio.Options{Transports: []transport.Transport{pt, wt}})
	if err != nil {
		log.Fatal(err)
	}

	websockets.SocketConnections(server)
	websockets.SocketEvents(server)
	websockets.ErrorConnections(server)

	exampleRoutes(mux)

	go server.Serve()
	defer server.Close()

	mux.Handle("/socket.io/", server)
	http.Handle("/", http.FileServer(http.Dir("./asset")))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4206"},
		AllowedMethods:   []string{"GET", "PUT", "OPTIONS", "POST", "DELETE"},
		AllowCredentials: true,
	})

	handler := c.Handler(mux)
	log.Println("Serving at localhost:8000...")
	log.Fatal(http.ListenAndServe(":8000", handler))
}
