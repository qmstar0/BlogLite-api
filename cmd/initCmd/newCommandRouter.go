package initCmd

import (
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/ThreeDotsLabs/watermill/message/router/middleware"
	"time"
)

var duration = time.Second * 10

func NewCommandRouter() *message.Router {
	newRouter, err := message.NewRouter(message.RouterConfig{CloseTimeout: duration}, nil)
	newRouter.AddMiddleware(middleware.Retry{
		MaxRetries:          3,
		InitialInterval:     time.Second,
		MaxInterval:         0,
		Multiplier:          0,
		MaxElapsedTime:      0,
		RandomizationFactor: 0,
		OnRetryHook:         nil,
		Logger:              nil,
	}.Middleware)
	if err != nil {
		panic(err)
	}
	return newRouter
}
