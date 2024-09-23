package adapter

import (
	"context"
	"fmt"
	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"github.com/ThreeDotsLabs/watermill/message/router/plugin"
	"github.com/qmstar0/BlogLite-api/internal/common/gopubsub"
	"github.com/qmstar0/BlogLite-api/pkg/logging"
	"time"
)

type Bus struct {
	router *message.Router
	message.Publisher
	sub message.Subscriber
}

func NewBus() *Bus {
	logger := gopubsub.NewLogger()
	router, _ := message.NewRouter(message.RouterConfig{CloseTimeout: time.Second * 5}, logger)
	router.AddPlugin(plugin.SignalsHandler)
	router.AddMiddleware(
		// 错误恢复
		middleware.Recoverer,
		// 忽略handler返回的错误并记录日志
		func(h message.HandlerFunc) message.HandlerFunc {
			return func(msg *message.Message) ([]*message.Message, error) {
				messages, err := h(msg)
				if err != nil {
					logger.Error("Message processing fails after retries", err, watermill.LogFields{"message_uuid": msg.UUID})
				}
				return messages, nil
			}
		},

		// handler重试
		middleware.Retry{
			MaxRetries:          5,
			InitialInterval:     time.Millisecond * 500,
			MaxInterval:         time.Second * 5,
			Multiplier:          1.5,
			MaxElapsedTime:      time.Second * 30,
			RandomizationFactor: 0.5,
			Logger:              logger,
		}.Middleware,
		//middleware.RandomFail(0.5),
	)
	pubsub := gopubsub.NewPubsub()
	return &Bus{
		router:    router,
		Publisher: pubsub,
		sub:       pubsub,
	}
}

func (b *Bus) Register(name string, handler gopubsub.Handler) {
	for _, topic := range handler.Topics() {
		b.router.AddHandler(
			b.name(name, topic),
			topic,
			b.sub,
			"",
			b.Publisher,
			handler.Handle,
		)
	}
}

func (b *Bus) Run(ctx context.Context) {
	go func() {
		err := b.router.Run(ctx)
		if err != nil {
			logging.Logger().Fatal("message router startup failed", "err", err)
		}
	}()
}

func (b *Bus) RouterClose() {
	_ = b.router.Close()
}

func (b *Bus) name(name, topic string) string {
	return fmt.Sprintf("%s.(%s)", name, topic)
}
