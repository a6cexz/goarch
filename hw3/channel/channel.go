package channel

import "fmt"

// Subscribers map of Subscriber interface
type Subscribers map[string]Subscriber

// Channel struct
type Channel struct {
	subscribers Subscribers
}

// NewChannel creates new Channel
func NewChannel() *Channel {
	return &Channel{
		subscribers: Subscribers{},
	}
}

// Send sends message to all subsribers in the channel
func (ch *Channel) Send(msg string) {
	for _, sub := range ch.subscribers {
		sub.OnReceive(msg)
	}
}

// Subscribe adds new Subscriber to the channel
func (ch *Channel) Subscribe(sub Subscriber) {
	ch.subscribers[sub.GetID()] = sub
}

// UnSubscribe removes the given subscriber from the channel
func (ch *Channel) UnSubscribe(sub Subscriber) error {
	id := sub.GetID()
	if _, ok := ch.subscribers[id]; ok {
		delete(ch.subscribers, id)
		return nil
	}
	return fmt.Errorf("can't find user %s", id)
}

// UnSubscribeAll removes all subscribers
func (ch *Channel) UnSubscribeAll() error {
	for _, sub := range ch.subscribers {
		return ch.UnSubscribe(sub)
	}
	return nil
}
