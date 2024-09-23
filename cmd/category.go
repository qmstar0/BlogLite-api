package cmd

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/qmstar0/BlogLite-api/internal/categories/ports"
	"github.com/qmstar0/BlogLite-api/internal/categories/service"
	"github.com/qmstar0/BlogLite-api/internal/common/auth"
)

func RegisterCategoryServer(ctx context.Context, group gin.IRouter) {
	application := service.NewApplication(ctx)
	server := ports.NewHttpServer(application)
	ports.RegisterHandlersWithOptions(group, server, ports.GinServerOptions{
		Middlewares: []ports.MiddlewareFunc{auth.FilterUnloggedUsersMiddleware()},
	})
}
