package main

import (
	"blog/infra/config"
	"blog/infra/repository/database/redis"
	"blog/router/index"
	shutdown "blog/utils"
)

func init() {
	shutdown.WaitForShutdown().Add(func() { redis.Close() })
}

func main() {
	e := index.Router()
	if err := e.Run(config.Conf.Service.Addr); err != nil {
		panic(err)
	}
}
