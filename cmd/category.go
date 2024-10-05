package cmd

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/qmstar0/BlogLite-api/internal/categories/ports"
	"github.com/qmstar0/BlogLite-api/internal/categories/service"
)

func RegisterCategoryServer(ctx context.Context, group gin.IRouter) {
	application := service.NewApplication(ctx)
	server := ports.NewHttpServer(application)
	ports.RegisterHandlers(group, server)
}
