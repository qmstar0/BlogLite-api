package main

import (
	"go-blog-ddd/config"
	"go-blog-ddd/internal/adapter/postgresql"
	"go-blog-ddd/internal/adapter/utils"
	"go-blog-ddd/internal/application"
	"go-blog-ddd/internal/ports/httpserver"
)

func init() {
	config.Init()
	closeFn := postgresql.Init()
	utils.OnShutdown(closeFn)
}

func main() {

	app := application.NewApp()
	server := httpserver.NewHttpServer(app)

	httpserver.RunHttpServer(config.Cfg.App.Addr, server)
}
