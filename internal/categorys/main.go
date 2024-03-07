package main

import (
	"blog/pkg/env"
	"blog/pkg/postgresql"
	"categorys/adapter"
	"categorys/ports"
	"common/server"
	"github.com/go-chi/chi/v5"
)

func init() {
	env.Load()
	postgresql.Init()
}
func main() {
	app := adapter.NewApp()

	server.RunHttpServer(":3000", func(router chi.Router) {
		router.Route("/api", func(r chi.Router) {
			ports.HandlerFromMuxWithBaseURL(ports.NewHttpServer(app), r, "/categorys")
		})
	})
}
