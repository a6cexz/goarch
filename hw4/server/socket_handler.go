package server

import (
	"fmt"
	"log"
	"net/http"
	"realtimechat/messages"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

// SocketHandler api
func (serv *Server) SocketHandler(w http.ResponseWriter, r *http.Request) {
	ws, err := serv.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatalf("websocket err: %v", err)
	}

	go func() {
		for {
			<-time.After(5 * time.Second)
			msg := messages.Message{
				Type: messages.Ping,
			}
			if err := ws.WriteJSON(msg); err != nil {
				log.Printf("ws send ping err: %v", err)
			}
		}
	}()

	serv.submutex.Lock()

	id := uuid.New().String()
	user := NewChatUser(id, ws)
	mainChannel := serv.publisher.GetChannel("#main")
	mainChannel.Subscribe(user)
	serv.submutex.Unlock()

	for {
		msg := messages.Message{}
		if err := ws.ReadJSON(&msg); err != nil {
			if !websocket.IsCloseError(err, 1001) {
				log.Fatalf("ws msg read err: %v", err)
			}
			break
		}

		if msg.IsPong() {
			continue
		}

		if msg.IsText() {
			fmt.Println(msg.Data)
			serv.submutex.Lock()
			if err := serv.publisher.Send(msg.Data); err != nil {
				log.Fatalf("ws msg subs err: %v", err)
			}
			serv.submutex.Unlock()
		}
	}

	fmt.Println("CLSOED")
	defer func() {
		serv.submutex.Lock()
		if err := serv.publisher.DeleteChannel(id); err != nil {
			log.Fatalf("Can not delete channel: %v", err)
		}
		serv.submutex.Unlock()
	}()
}
