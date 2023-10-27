package validate

import (
	"github.com/gin-gonic/gin"
)

type AuthV struct {
}

func (a AuthV) Validate() gin.HandlerFunc {
	return func(c *gin.Context) {
		var (
		//apiC = response.Api{C: c}
		)
		c.Next()
	}
}
