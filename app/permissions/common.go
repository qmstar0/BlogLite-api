package permissions

import "github.com/gin-gonic/gin"

type P interface {
	Permission() gin.HandlerFunc
}

type Permission struct {
	NewAuthP P
}

func NewPermission() *Permission {
	return &Permission{
		NewAuthP: AuthP{},
	}
}
