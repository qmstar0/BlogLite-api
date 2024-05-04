package main

import (
	"go-blog-ddd/config"
	"go-blog-ddd/internal/adapter/postgresql"
	"go-blog-ddd/internal/adapter/utils/shutdown"
	"go-blog-ddd/internal/application"
	"go-blog-ddd/internal/ports/httpserver"
)

func init() {
	config.Init()
	fn := postgresql.Init()
	shutdown.OnShutdown(fn)
}

func main() {
	app := application.NewApp()
	server := httpserver.NewHttpServer(app)
	httpserver.RunHttpServer(":3000", server)
}
