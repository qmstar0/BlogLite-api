package auth

import (
	"github.com/gin-gonic/gin"
	"github.com/qmstar0/BlogLite-api/config"
	"github.com/qmstar0/BlogLite-api/internal/common/constant"
	"github.com/qmstar0/BlogLite-api/internal/common/e"
	"github.com/qmstar0/BlogLite-api/internal/common/server/httpresponse"
	"github.com/qmstar0/BlogLite-api/pkg/utils"
)

type GetAuthJSONParams struct {
	Password string `json:"password"`
}

func GetAuth(c *gin.Context) {
	req, err := utils.BindJSON[GetAuthJSONParams](c)
	if err != nil {
		httpresponse.Error(c, err)
		return
	}

	if req.Password != config.Cfg.AuthAdminPassword {
		httpresponse.Error(c, e.PWDError)
		return
	}

	sign, err := Sign(NewUserCliaims(Admin.ID, Admin.Type, Admin.Name, constant.DefaultJWTAuthDuration))
	if err != nil {
		httpresponse.Error(c, e.InternalServiceError(err.Error()))
		return
	}

	c.Header("X-Auth-Token", sign)
	httpresponse.OK(c)
}
