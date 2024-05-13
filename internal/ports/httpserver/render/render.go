package render

import (
	"github.com/gin-gonic/gin"
	"go-blog-ddd/config"
	"go-blog-ddd/internal/adapter/e"
)

var errorToResponse func(error) Response

func init() {
	if config.Cfg.Build.Release {
		errorToResponse = releaseModeErrorResponse
	} else {
		errorToResponse = debugModeErrorResponse
	}
}

func Error(c *gin.Context, err error) {
	c.JSON(200, errorToResponse(err))
}

func Success(c *gin.Context) {
	c.JSON(200, e.Successed)
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
}
