package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func AuthorizerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authToken := getAuthTokenFromHeader(c)
		if authToken != "" {
			if authVerified, err := Srv.VaildateAuth(authToken); err == nil {
				c.Set("userId", authVerified.Uid)
				c.Set("email", authVerified.Email)
				c.Set("role", authVerified.Role)
			}
		}
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
