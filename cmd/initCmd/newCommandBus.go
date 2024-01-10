package initCmd

import (
	"fmt"
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
)

func NewCommandBus(pubsub PubSub, marshaler cqrs.CommandEventMarshaler) *cqrs.CommandBus {
	cmdBus, err := cqrs.NewCommandBusWithConfig(pubsub, cqrs.CommandBusConfig{
		GeneratePublishTopic: func(params cqrs.CommandBusGeneratePublishTopicParams) (string, error) {
			fmt.Println("发布", params.CommandName)
			return params.CommandName, nil
		},
		OnSend:    nil,
		Marshaler: marshaler,
		Logger:    nil,
	})
	if err != nil {
		panic(err)
	}
	return cmdBus

}
