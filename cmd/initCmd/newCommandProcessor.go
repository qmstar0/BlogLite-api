package initCmd

import (
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
	"github.com/ThreeDotsLabs/watermill/message"
)

func NewCommandProcessor(pubsub PubSub, router *message.Router, marshaler cqrs.CommandEventMarshaler) *cqrs.CommandProcessor {
	cmdProcessor, err := cqrs.NewCommandProcessorWithConfig(router, cqrs.CommandProcessorConfig{
		GenerateSubscribeTopic: func(params cqrs.CommandProcessorGenerateSubscribeTopicParams) (string, error) {
			return params.CommandName, nil
		},
		SubscriberConstructor: func(params cqrs.CommandProcessorSubscriberConstructorParams) (message.Subscriber, error) {
			return pubsub, nil
		},
		OnHandle:                 nil,
		Marshaler:                marshaler,
		Logger:                   nil,
		AckCommandHandlingErrors: true,
	})
	if err != nil {
		panic(err)
	}
	return cmdProcessor
}
