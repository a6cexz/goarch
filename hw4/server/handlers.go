package server

import (
	"net/http"
)

// ApplyHandlers setups chat handlers
func (serv *Server) ApplyHandlers() {
	serv.router.Handle("/*", http.FileServer(http.Dir("./web")))
	serv.router.Get("/socket", serv.SocketHandler)
}
