package CQRS

import (
	"context"
	"github.com/ThreeDotsLabs/watermill/message"
)

type CommandBus struct {
	publisher message.Publisher
	marshal   CommandEventMarshal
}

func NewCommandBus(publisher message.Publisher, marshal CommandEventMarshal) *CommandBus {
	return &CommandBus{publisher: publisher, marshal: marshal}
}

func (c *CommandBus) Send(ctx context.Context, cmd TypeCommand) error {
	msg, err := c.marshal.Marshal(cmd)
	if err != nil {
		return err
	}
	msg.SetContext(ctx)
	return c.publisher.Publish(
		cmd.Topic(),
		msg,
	)
}
