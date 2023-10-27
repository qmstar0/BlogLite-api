package validate

import (
	"blog/app/dto"
	"blog/app/response"
	"blog/infra/e"
	"github.com/gin-gonic/gin"
)

type CaptchaV struct {
}

func (c CaptchaV) Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
			apiC = response.Api{C: c}
			req  = dto.Captcha{}
		)
		if err := c.ShouldBind(&req); err != nil {
			apiC.ValidateFailResp(err)
			return
		}

		if err := validate.Struct(&req); err != nil {
			apiC.ValidateFailResp(e.NewError(e.InvalidParam, err))
			return
		}
		c.Set("captcha", req)
		c.Next()
	}
}
