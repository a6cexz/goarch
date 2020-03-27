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

func getSubsciberSendFunc(ws *websocket.Conn) Subscriber {
	return func(msg string) error {
		m := messages.Message{
			Type: messages.Text,
			Data: msg,
		}
		if err := ws.WriteJSON(m); err != nil {
			log.Printf("ws msg fetch err: %v", err)
		}
		return nil
	}
}

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

	id := uuid.New().String()
	serv.submutex.Lock()
	serv.subscribers[id] = getSubsciberSendFunc(ws)
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
			for _, sub := range serv.subscribers {
				if err := sub(msg.Data); err != nil {
					log.Fatalf("ws msg subs err: %v", err)
				}
			}
			serv.submutex.Unlock()
		}
	}

	fmt.Println("CLSOED")
	defer func() {
		serv.submutex.Lock()
		delete(serv.subscribers, id)
		serv.submutex.Unlock()
	}()
}
