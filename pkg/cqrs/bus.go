package cqrs

import (
	"context"
	"fmt"
	"github.com/qmstar0/eio"
	"time"
)

type Bus struct {
	pub         eio.Publisher
	sub         eio.Subscriber
	codec       eio.Codec
	newIdFn     func() string
	onHandleErr func(topic string)
}

func NewBus(pub eio.Publisher, sub eio.Subscriber, codec eio.Codec, newId func() string) *Bus {
	return &Bus{
		pub:     pub,
		sub:     sub,
		codec:   codec,
		newIdFn: newId,
	}
}
func (b Bus) Publish(c context.Context, v any) error {
	topic := getStructName(v)

	message, err := b.getMessage(v)
	if err != nil {
		return err
	}
	return b.pub.Publish(c, topic, message)
}

func (b Bus) AddHandler(handler ...Handler) {
	for i := range handler {
		b.addHandler(handler[i])
	}
}

func (b Bus) addHandler(handler Handler) {
	topic := handler.Topic()
	ctx := context.TODO()
	messageCh, err := b.sub.Subscribe(ctx, topic)
	if err != nil {
		panic(err)
	}

	go b.listenMessage(messageCh, b.getHandleMessageFn(ctx, handler))
}

func (b Bus) listenMessage(msgCh <-chan eio.Message, handlerFn func(message eio.Message) error) {
	for message := range msgCh {
		err := handlerFn(message)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func (b Bus) getHandleMessageFn(ctx context.Context, handler Handler) func(message eio.Message) error {
	var (
		err       error
		id        string
		timestamp time.Time
		data      = handler.SubscribeTo()
	)
	return func(message eio.Message) error {
		err = b.codec.Decode(message, &id, &timestamp, data)
		if err != nil {
			return fmt.Errorf("err on decoding: %w", err)
		}
		toCtx := setMessageInfoToCtx(ctx, id, timestamp)
		err = handler.Handle(toCtx, data)
		if err != nil {
			return fmt.Errorf("err on handle: %w", err)
		}
		return nil
	}
}

func (b Bus) getMessage(v any) (eio.Message, error) {
	id := b.newIdFn()
	timestamp := time.Now()
	return b.codec.Encode(id, timestamp, v)
}

func setMessageInfoToCtx(ctx context.Context, id string, timestamp time.Time) context.Context {
	return context.WithValue(
		context.WithValue(ctx, ctxKeyId, id),
		ctxKeyTimestamp, timestamp)
}
