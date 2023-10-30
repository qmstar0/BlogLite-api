package validate

import (
	"blog/app/dto"
	"blog/app/response"
	"blog/infra/e"
	"github.com/gin-gonic/gin"
)

type UserV struct {
}

func (u UserV) Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			apiC = response.Api{C: c}
			req  = dto.UserStore{}
		)
		if err := c.ShouldBind(&req); err != nil {
			apiC.Response(err)
			return
		}
		if err := validate.Struct(&req); err != nil {
			apiC.Response(e.NewError(e.InvalidParam, err))
			return
		}
		c.Set("store", req)
		c.Next()
	}
}
