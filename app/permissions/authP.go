package permissions

import (
	"blog/app/response"
	"blog/infra/e"
	"github.com/gin-gonic/gin"
)

type AuthP struct {
}

func (a AuthP) Permission() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			apiC = response.Api{C: c}
		)
		userId := c.GetString("userId")
		if userId == "" {
			apiC.Response(e.NewError(e.LoginRequired, nil))
			return
		}
		c.Next()
	}
}
