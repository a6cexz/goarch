package messages

// MessageType string
type MessageType string

const (
	// Ping message type
	Ping MessageType = "ping"

	// Pong message type
	Pong MessageType = "pong"

	// Text message type
	Text MessageType = "message"
)

// Message data
type Message struct {
	Type MessageType `json:"type"`
	Data string      `json:"data,omitempty"`
}

// NewPing creates new ping message
func NewPing() *Message {
	msg := &Message{
		Type: Ping,
	}
	return msg
}

// NewPong creates new pong message
func NewPong() *Message {
	msg := &Message{
		Type: Pong,
	}
	return msg
}

// NewText creates new text message
func NewText(data string) *Message {
	msg := &Message{
		Type: Text,
		Data: data,
	}
	return msg
}

// IsPing returns true if message has Ping type
func (msg *Message) IsPing() bool {
	return msg.Type == Ping
}

// IsPong returns true if message has Pong type
func (msg *Message) IsPong() bool {
	return msg.Type == Pong
}

// IsText returns true if message has Text type
func (msg *Message) IsText() bool {
	return msg.Type == Text
}
