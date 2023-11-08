package permissions

import (
	"blog/domain/users/valueobject"
	"github.com/gin-gonic/gin"
)

type PublishP struct {
}

func (p PublishP) Check(c *gin.Context) bool {
	userRole := c.GetUint("role")
	if userRole == 0 {
		return false
	}
	role := valueobject.NewUserRole(userRole)
	if !role.IsPublisher() {
		return false
	}
	return true

}
