package main

import (
	"blog/pkg/cqrs"
	"blog/pkg/env"
	"blog/pkg/postgresql"
	"categorys/adapter"
	"categorys/ports"
	"common/server"
	"github.com/go-chi/chi/v5"
	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
	"github.com/qmstar0/eio"
	"github.com/qmstar0/eio-redis/redispubsub"
)

func init() {
	env.Load()
	postgresql.InitDB()
}
func main() {
	redisCli := redis.NewClient(&redis.Options{
		Addr:     "192.168.1.10:6379", //Testing in LAN environment
		Password: "",
		DB:       0,
	})

	pub := redispubsub.NewRedisPublisher(redisCli)
	defer pub.Close()
	sub := redispubsub.NewRedisSubscriber(redisCli)
	defer sub.Close()

	bus := cqrs.NewBus(pub, sub, eio.NewJSONCodec(), func() string {
		return uuid.New().String()
	})

	app := adapter.NewApp(bus)

	server.RunHttpServer(":3000", func(router chi.Router) {
		router.Route("/api", func(r chi.Router) {
			//ports.HandlerFromMuxWithBaseURL(ports.NewHttpServer(app), r, "/categorys")
			ports.HandlerWithOptions(ports.NewHttpServer(app), ports.ChiServerOptions{
				BaseURL:          "/categorys",
				BaseRouter:       r,
				Middlewares:      append([]ports.MiddlewareFunc(nil), server.AuthMiddleware()),
				ErrorHandlerFunc: nil,
			})
		})

	})

}
