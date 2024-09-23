package cmd

import (
	"context"
	"flag"
	"github.com/gin-gonic/gin"
	"github.com/qmstar0/BlogLite-api/config"
	"github.com/qmstar0/BlogLite-api/internal/common/server"
	"github.com/qmstar0/BlogLite-api/pkg/logging"
	"github.com/qmstar0/BlogLite-api/pkg/postgresql"
	"github.com/qmstar0/shutdown"
)

const defaultConfigPath = "./config.toml"

func InitDatabase() {
	closeFn := postgresql.Init(postgresql.PostgresDSN(config.Cfg.DatabaseDNS))
	shutdown.RegisterTasks(func() { _ = closeFn() })
}

func InitLogger() {
	logging.Init(config.Cfg.Mode == "debug")
}

func Run() {
	var configPath string
	flag.StringVar(&configPath, "config", defaultConfigPath, "输入配置文件的文件路径")
	flag.Parse()

	config.Init(configPath)
	InitLogger()
	InitDatabase()

	ctx := context.Background()
	server.RunHTTPServer(config.Cfg.PORT, func(router gin.IRouter) {
		RegisterArticleServer(ctx, router)
		RegisterCategoryServer(ctx, router)
		RegisterAuthServer(ctx, router)
	})
}
