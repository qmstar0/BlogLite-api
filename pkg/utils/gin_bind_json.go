package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/qmstar0/BlogLite-api/internal/common/e"
)

func BindJSON[T any](c *gin.Context) (T, error) {
	var req T
	if err := c.ShouldBindJSON(&req); err != nil {
		return req, e.InvalidParametersError(err.Error())
	}
	return req, nil
}
