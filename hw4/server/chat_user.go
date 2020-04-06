package server

import (
	"log"
	"realtimechat/channel"
	"realtimechat/messages"

	"github.com/gorilla/websocket"
)

// ChatUser service user
type ChatUser struct {
	channel.User
	ws *websocket.Conn
}

// NewChatUser creates new chat service user
func NewChatUser(username string, ws *websocket.Conn) *ChatUser {
	return &ChatUser{
		User: channel.User{
			Username: username,
		},
		ws: ws,
	}
}

// OnReceive hanles new received message
func (u *ChatUser) OnReceive(msg string) {
	m := messages.Message{
		Type: messages.Text,
		Data: msg,
	}
	if err := u.ws.WriteJSON(m); err != nil {
		log.Printf("ws msg fetch err: %v", err)
	}
}
