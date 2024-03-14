package main

import (
	_ "blog/pkg/env"
	"categorys/adapter"
	"categorys/ports"
	"common/server"
	"github.com/go-chi/chi/v5"
)

func main() {
	app := adapter.NewApp()
	server.RunHttpServer(":3000", func(r chi.Router) {
		ports.HandlerWithOptions(ports.NewHttpServer(app), ports.ChiServerOptions{
			BaseRouter: r,
		})
	})
}
