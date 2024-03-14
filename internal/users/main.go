package main

import (
	_ "blog/pkg/env"
	"common/server"
	"github.com/go-chi/chi/v5"
	"users/adapters"
	"users/ports"
)

func main() {
	app := adapters.NewApp()
	server.RunHttpServer(":3000", func(r chi.Router) {
		ports.HandlerWithOptions(ports.NewHttpServer(app), ports.ChiServerOptions{
			BaseRouter: r,
		})
	})
}
