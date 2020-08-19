package routes

import (
	"net/http"
)

func exampleRoutes(server *http.ServeMux) {


	server.HandleFunc("/hola", func(writer http.ResponseWriter, request *http.Request) {
		println("sadasd ")
	})
}
