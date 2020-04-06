package channel

import (
	"fmt"
)

// Channels map of Channel
type Channels map[string]*Channel

// Publisher struct
type Publisher struct {
	channels Channels
}

// NewPublisher creates new Publisher
func NewPublisher() *Publisher {
	return &Publisher{
		channels: Channels{},
	}
}

// AddChannel adds new channel to the publisher
func (p *Publisher) AddChannel(name string, channel *Channel) {
	p.channels[name] = channel
}

// GetChannels returns all channels
func (p *Publisher) GetChannels() []*Channel {
	r := make([]*Channel, 0, len(p.channels))
	for _, ch := range p.channels {
		r = append(r, ch)
	}
	return r
}

// GetChannel returns channel by name
func (p *Publisher) GetChannel(name string) *Channel {
	ch, ok := p.channels[name]
	if ok {
		return ch
	}
	return nil
}

// DeleteChannel deletes channel with the given name
func (p *Publisher) DeleteChannel(name string) error {
	ch, ok := p.channels[name]
	if !ok {
		return fmt.Errorf("Can not delete %v channel, because it does not exist", name)
	}
	ch.UnSubscribeAll()
	delete(p.channels, name)
	return nil
}

// Send send message to the given list of channels or to all channels if list is empty
func (p *Publisher) Send(msg string, channels ...string) error {
	if len(channels) == 0 {
		for _, ch := range p.channels {
			ch.Send(msg)
		}
	} else {
		for _, ch := range channels {
			channel, ok := p.channels[ch]
			if !ok {
				return fmt.Errorf("channel %s can't be found", ch)
			}
			channel.Send(msg)
		}
	}
	return nil
}
