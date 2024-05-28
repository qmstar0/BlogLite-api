package httpserver

import (
	"github.com/gin-gonic/gin"
	"github.com/qmstar0/nightsky-api/internal/apps/service"
	"github.com/qmstar0/nightsky-api/internal/pkg/e"
	"github.com/qmstar0/nightsky-api/ports/httpserver/render"

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
