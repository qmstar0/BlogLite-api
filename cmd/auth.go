package cmd

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/qmstar0/BlogLite-api/internal/common/auth"
)

func RegisterAuthServer(_ context.Context, group gin.IRouter) {
	group.POST("/authentication", auth.GetAuth)
}
