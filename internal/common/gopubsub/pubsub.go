package gopubsub

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/pubsub/gochannel"
	"sync"
)

type Handler interface {
	Topics() []string
	Handle(msg *message.Message) ([]*message.Message, error)
}

type Pubsub interface {
	message.Publisher
	message.Subscriber
}

var (
	pubsub     Pubsub
	pubsubOnce sync.Once
)

func NewPubsub() Pubsub {
	pubsubOnce.Do(func() {
		pubsub = gochannel.NewGoChannel(gochannel.Config{}, NewLogger())
	})
	return pubsub
}
