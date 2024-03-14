package main

import (
	"blog/pkg/env"
	"blog/pkg/mongodb"
	"blog/pkg/rediscache"
	"blog/pkg/shutdown"
	"common/auth"
	"common/server"
	"context"
	"posts/adapter"
	"posts/ports"
)

func init() {
	env.Load()

	dbCloseFn := mongodb.Init()
	shutdown.OnShutdown(func() error { return dbCloseFn(context.Background()) })

	cacheCloseFn := rediscache.Init()
	shutdown.OnShutdown(cacheCloseFn)
}
func main() {

	app := adapter.NewApp()

	server.RunHttpServer(":3000", func(r chi.Router) {

		ports.HandlerWithOptions(ports.NewHttpServer(app), ports.ChiServerOptions{
			BaseRouter:  r,
			Middlewares: append([]ports.MiddlewareFunc(nil), auth.AuthMiddleware()),
		})

	})
}
