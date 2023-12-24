package common

import "github.com/ThreeDotsLabs/watermill/message"

type CommandHandler interface {
	Handle(msg *message.Message) ([]*message.Message, error)
}
