package validate

import (
	"blog/app/dto"
	"blog/app/response"
	"blog/infra/e"
	"github.com/gin-gonic/gin"
)

type CateV struct {
}

func (c CateV) Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			apiC = response.Api{C: c}
			req  = dto.CateStore{}
		)
		if err := c.ShouldBind(&req); err != nil {
			apiC.ValidateFailResp(err)
			return
		}
		if err := validate.Struct(&req); err != nil {
			apiC.ValidateFailResp(e.NewError(e.InvalidParam, err))
			return
		}
		c.Set("store", req)
		c.Next()
	}
}
