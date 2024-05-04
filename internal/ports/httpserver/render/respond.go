package render

import (
	"errors"
	"github.com/gin-gonic/gin"
	"go-blog-ddd/internal/application/e"
)

var OK = func(data any) Response { return Response{StateCode: e.Successed, Data: data} }

type Response struct {
	e.StateCode `json:",inline"`
	Data        any `json:"data,omitempty"`
}

func Error(c *gin.Context, err error) {
	var result e.StateCode
	if !errors.As(err, &result) {
		result = e.NotImplemented(err.Error())
	}
	c.JSON(200, result)
}

func Success(c *gin.Context) {
	c.JSON(200, OK(nil))
}

func Respond(c *gin.Context, data any) {
	c.JSON(200, OK(data))
}
