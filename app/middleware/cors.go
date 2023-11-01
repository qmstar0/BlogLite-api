package middleware

import "github.com/gin-gonic/gin"

// CORSMiddleware 跨域配置
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")                                // 允许任何来源访问
		c.Header("Access-Control-Allow-Methods", "POST, GET, PUT, DELETE, OPTIONS") // 允许的HTTP方法
		c.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")     // 允许的请求标头
		c.Header("Access-Control-Allow-Credentials", "true")                        // 允许发送凭证，如Cookies
		c.Header("Access-Control-Max-Age", "3600")                                  // 预检请求的有效期，单位秒

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204) // 预检请求返回204 No Content
		} else {
			c.Next()
		}
	}
}
