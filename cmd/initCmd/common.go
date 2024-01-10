package initCmd

import "github.com/ThreeDotsLabs/watermill/message"

type PubSub interface {
	message.Publisher
	message.Subscriber
}
