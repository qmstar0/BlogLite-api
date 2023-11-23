package middleware

import (
	"blog/app/response"
	"blog/infra/e"
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthorizerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var apiC = response.Api{C: c}
		authToken := getAuthTokenFromHeader(c)
		if authToken == "" {
			apiC.Response(e.NewError(e.PermissionDenied, nil))
			return
		}
		authVerified, err := Srv.VaildateAuth(authToken)
		if err != nil {
			apiC.Response(e.NewError(e.PermissionDenied, err))
			return
		}
		c.Set("userId", authVerified["userId"])
		c.Set("email", authVerified["email"])
		c.Set("role", authVerified["role"])
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
