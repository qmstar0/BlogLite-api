package middleware

import (
	"blog/app/response"
	"blog/infra/e"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthenticateMiddleware() func(ctx *gin.Context) {
	return func(c *gin.Context) {
		var (
			err  error
			apiC = response.Api{C: c}
		)
		authToken := getAuthTokenFromHeader(c)
		if authToken == "" {
			apiC.Response(e.NewError(e.PermissionDenied, nil))
			return
		}
		authVerified, err := Srv.VaildateAuth(authToken)
		if err != nil {
			apiC.Response(e.NewError(e.TokenVerifyErr, err))
			return
		}
		c.Set("userId", authVerified.Uid)
		c.Set("email", authVerified.Email)
		c.Set("role", authVerified.Role)
		c.Next()
	}
}

func getAuthTokenFromHeader(c *gin.Context) string {
	//return c.GetHeader("x-auth-token")
	authStr := c.GetHeader("Authorization")
	if authStr == "" {
		return ""
	}
	return strings.Split(authStr, " ")[1]
}
