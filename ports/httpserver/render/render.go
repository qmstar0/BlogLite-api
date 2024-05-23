package render

import (
	"github.com/gin-gonic/gin"
	"go-blog-ddd/config"
	"go-blog-ddd/internal/pkg/e"
)

var errorToResponse func(error) Response

func init() {
	if config.Cfg.Release {
		errorToResponse = releaseModeErrorResponse
	} else {
		errorToResponse = debugModeErrorResponse
	}
}

func Error(c *gin.Context, err error) {
	c.JSON(200, errorToResponse(err))
	c.Abort()
}

func Success(c *gin.Context) {
	c.JSON(200, e.Successed)
	c.Abort()
}

func Respond(c *gin.Context, err error, data any) {
	if err != nil {
		c.JSON(200, errorToResponse(err))
	} else {
		c.JSON(200, Response{
			StateCode: e.Successed,
			Data:      data,
		})
	}
	c.Abort()
}
