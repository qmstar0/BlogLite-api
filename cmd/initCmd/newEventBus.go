package initCmd

import (
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
)

func NewEventBus(pubsub PubSub, marshaler cqrs.CommandEventMarshaler) *cqrs.EventBus {
	eventBus, err := cqrs.NewEventBusWithConfig(pubsub, cqrs.EventBusConfig{
		GeneratePublishTopic: func(params cqrs.GenerateEventPublishTopicParams) (string, error) {
			return params.EventName, nil
		},

		OnPublish: nil,

		Marshaler: marshaler,

		Logger: nil,
	})
	if err != nil {
		panic(err)
	}
	return eventBus
}
