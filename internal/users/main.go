package users

import (
	"blog/pkg/env"
	"blog/pkg/mongodb"
	"blog/pkg/rediscache"
	"blog/pkg/shutdown"
	"common/server"
	"context"
	"github.com/go-chi/chi/v5"
	"users/adapters"
	"users/ports"
)

func init() {
	env.Load()

	dbCloseFn := mongodb.Init()
	shutdown.OnShutdown(func() error { return dbCloseFn(context.Background()) })

	cacheCloseFn := rediscache.Init()
	shutdown.OnShutdown(cacheCloseFn)
}
func main() {
	app := adapters.NewApp()
	server.RunHttpServer(":3000", func(r chi.Router) {
		ports.HandlerWithOptions(ports.NewHttpServer(app), ports.ChiServerOptions{
			BaseRouter: r,
		})
	})
}
