package httpserver

import (
	"github.com/gin-gonic/gin"
	"go-blog-ddd/internal/apps/service"
	"go-blog-ddd/internal/pkg/e"
	"go-blog-ddd/ports/httpserver/render"
	"strings"
)

func AuthMiddleware(service *service.AdminAuthenticationService) MiddlewareFunc {
	return func(c *gin.Context) {
		_, needAuth := c.Get(BearerScopes)
		if needAuth {
			token, err := getAuthorizationToken(c)
			if err != nil {
				render.Error(c, err)
				return
			}
			err = service.Verify(token)
			if err != nil {
				render.Error(c, err)
				return
			}
		}
		c.Next()
	}
}

func getAuthorizationToken(c *gin.Context) (string, error) {
	Authorization := c.GetHeader("Authorization")
	if len(Authorization) < 7 {
		return "", e.AErrUnauthortion
	}
	splitAuthorization := strings.Split(Authorization, " ")
	if len(splitAuthorization) != 2 {
		return "", e.AErrUnauthortion
	}
	return splitAuthorization[1], nil
}
