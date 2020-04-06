package channel

// Subscriber interface
type Subscriber interface {
	OnReceive(msg string)
	GetID() string
}

// SubscriberDefault default struct
type SubscriberDefault struct{}

// OnReceive called when new message is recevied
func (SubscriberDefault) OnReceive(string) {
	panic("not implemented")
}

// GetID returns subscriber ID
func (SubscriberDefault) GetID(string) {
	panic("not implemented")
}
