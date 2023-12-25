package CQRS

import (
	"github.com/ThreeDotsLabs/watermill/message"
)

type EventBus struct {
	publisher message.Publisher
	marshal   CommandEventMarshal
}

func NewEventBus(publisher message.Publisher, marshal CommandEventMarshal) *EventBus {
	return &EventBus{publisher: publisher, marshal: marshal}
}

func (c *EventBus) Publish(event TypeEvent) error {
	newMessage, err := c.marshal.Marshal(event)
	if err != nil {
		return err
	}
	return c.publisher.Publish(
		event.Topic(),
		newMessage,
	)
}
