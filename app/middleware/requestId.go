package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"strconv"
	"time"
)

func RequestIDMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestID string
		beginTime := strconv.FormatInt(time.Now().UnixNano(), 10)
		requestID = c.Request.Header.Get("Set-Request-Id")

		if requestID == "" {
			s := uuid.New()
			requestID = s.String()
		}

		c.Writer.Header().Set("X-Begin-Time", beginTime)
		c.Writer.Header().Set("X-Request-Id", requestID)
		c.Next()
	}
}
