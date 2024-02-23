package cqrs_test

import (
	"blog/pkg/cqrs"
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/qmstar0/eio"
	"github.com/qmstar0/eio-redis/redispubsub"
	"testing"
	"time"
)

type Test struct {
	Name string
	Age  int
}

func TestNewBus(t *testing.T) {

	redisCli := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.10:6379", //Testing in LAN environment
		Password: "",
		DB:       0,
	})

	var (
		ctx, cc = context.WithTimeout(context.Background(), time.Second*5)
	)
	defer cc()

	pub := redispubsub.NewRedisPublisher(redisCli)
	defer pub.Close()
	sub := redispubsub.NewRedisSubscriber(redisCli)
	defer sub.Close()

	bus := cqrs.NewBus(pub, sub, eio.NewJSONCodec(), func() string {
		return uuid.New().String()
	})

	bus.AddHandler(cqrs.NewHandler[Test](func(ctx context.Context, v *Test) error {
		t.Log("run handler", v, cqrs.GetIdFromCtx(ctx), cqrs.GetTimestampFromCtx(ctx))
		return nil
	}))

	time.AfterFunc(time.Second*3, func() {
		sub.Close()
		t.Log("sub closed")
	})

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			default:
				err := bus.Publish(ctx, Test{
					Name: "QMstar",
					Age:  20,
				})
				if err != nil {
					t.Error(err)
				}
				time.Sleep(time.Millisecond * 300)
			}
		}
	}()
	<-ctx.Done()
}
