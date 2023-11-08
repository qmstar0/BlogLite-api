package permissions

import (
	"blog/app/response"
	"blog/infra/e"
	"github.com/gin-gonic/gin"
)

type P interface {
	Check(c *gin.Context) bool
}

type Permissioner struct {
	AuthP    P
	PublishP P
}

func NewPermissioner() *Permissioner {
	return &Permissioner{
		AuthP:    AuthP{},
		PublishP: PublishP{},
	}
}

func (P Permissioner) Permission(p P) gin.HandlerFunc {
	return func(context *gin.Context) {
		if !p.Check(context) {
			var apiC = response.Api{C: context}
			apiC.Response(e.NewError(e.PermissionDenied, nil))
		}
		context.Next()
	}
}
