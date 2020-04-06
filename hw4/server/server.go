package server

import (
	"net/http"
	"realtimechat/channel"
	"sync"

	"github.com/go-chi/chi"
	"github.com/gorilla/websocket"
)

// Server struct
type Server struct {
	router   *chi.Mux
	upgrader *websocket.Upgrader

	submutex  *sync.Mutex
	publisher *channel.Publisher
}

// New creates new server
func New() *Server {
	router := chi.NewRouter()

	upgrader := &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	serv := &Server{
		router:    router,
		upgrader:  upgrader,
		submutex:  &sync.Mutex{},
		publisher: channel.NewPublisher(),
	}

	mainChannel := channel.NewChannel()
	serv.publisher.AddChannel("#main", mainChannel)

	serv.ApplyHandlers()
	return serv
}

// Start starts server
func (serv *Server) Start() error {
	return http.ListenAndServe(":8085", serv.router)
}
