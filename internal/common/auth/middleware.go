package auth

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/qmstar0/BlogLite-api/internal/common/e"
	"strings"
)

func AuthMiddleware() func(*gin.Context) {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if authHeader == tokenString {
			c.Next()
			return
		}

		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), userTokenContextKey, tokenString))
		c.Next()
	}
}

const userTokenContextKey = "__userTokenContextKey"

func GetUserFromContext(ctx context.Context) (*UserClaims, error) {
	userTokenStr, exits := ctx.Value(userTokenContextKey).(string)
	if !exits {
		return nil, e.UnauthorisedError("用户未登录")
	}
	user, err := Verify(userTokenStr)
	if err != nil {
		return nil, e.UnauthorisedError(err.Error())
	}
	return user, nil
}
