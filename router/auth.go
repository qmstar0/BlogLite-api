package router

import "github.com/gin-gonic/gin"

type Auth interface {
	AuthRegister(c *gin.Context)
	AuthLogin(c *gin.Context)
	AuthResetPwd(c *gin.Context)
}

type User interface {
	Login(c *gin.Context)
	Register(c *gin.Context)
	Update(c *gin.Context)
	ResetPwd(c *gin.Context)
}
