package eventbus_test

import (
	"context"
	"go-blog-ddd/internal/adapter/eventbus"
	"testing"
	"time"
)

func TestName(t *testing.T) {
	bus := eventbus.NewEventBus(20)

	timeout, cancelFunc := context.WithTimeout(context.Background(), time.Second*3)
	defer cancelFunc()

	go func() {
		for {
			bus.Publish(events.EventQueue{"test"})
			time.Sleep(time.Millisecond * 50)
		}
	}()

	for evt := range bus.Channel(timeout) {
		t.Log("handle", evt)
	}
}
