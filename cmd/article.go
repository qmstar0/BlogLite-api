package cmd

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/qmstar0/BlogLite-api/internal/articles/ports"
	"github.com/qmstar0/BlogLite-api/internal/articles/service"
	"github.com/qmstar0/BlogLite-api/internal/common/auth"
)

func RegisterArticleServer(ctx context.Context, group gin.IRouter) {
	application := service.NewApplication(ctx)
	server := ports.NewHttpServer(application)
	ports.RegisterHandlersWithOptions(group, server, ports.GinServerOptions{
		Middlewares: []ports.MiddlewareFunc{auth.FilterUnloggedUsersMiddleware()},
	})
}
