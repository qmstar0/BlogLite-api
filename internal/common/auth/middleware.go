package auth

import (
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/qmstar0/BlogLite-api/internal/common/e"
	"github.com/qmstar0/BlogLite-api/internal/common/server/httpresponse"
	"strings"
)

func FilterUnloggedUsersMiddleware() func(*gin.Context) {
	return func(c *gin.Context) {
		if c.Request.Method != "GET" {
			user, err := GetUserFromContext(c.Request.Context())
			if err != nil {
				httpresponse.Error(c, err)
				c.Abort()
				return
			}

			if user.Type != "admin" {
				httpresponse.Error(c, errors.New("没有权限"))
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

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
