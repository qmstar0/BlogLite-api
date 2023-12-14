package handlers

import (
	"blog/app/response"
	"blog/domain/users"
	"blog/domain/users/valueobject"
	"blog/infra/config"
	"blog/infra/e"
	"blog/router"
	"blog/utils"
	"github.com/gin-gonic/gin"
)

type Auth struct {
}

func NewAuth() router.Auth {
	return &Auth{}
}
func (a Auth) AuthRegister(c *gin.Context) {
	var (
		err  error
		apiC = response.Api{C: c}
	)
	email := c.Query("email")
	newEmail, err := valueobject.NewEmail(email)
	if email == "" {
		apiC.Response(e.NewError(e.InvalidParam, err))
		return
	}
	if _, err = Dao.GetUser(c, &users.User{Email: newEmail}); err == nil {
		apiC.Response(e.NewError(e.UserDuplicateCreationErr, nil))
		return
	}
	codeNum := utils.RandomNum(config.Conf.User.CaptchaLength)
	token, err := Srv.GenCaptchaToken(email, codeNum)
	if err != nil {
		apiC.Response(err)
		return
	}
	if err = Srv.SendCaptchaEmail(email, codeNum); err != nil {
		apiC.Response(err)
		return
	}
	resp := make(map[string]any)
	resp["token"] = token
	resp["expTime"] = config.Conf.User.JwtCaptchaTokenLifeSec
	apiC.Success(resp)
}

func (a Auth) AuthLogin(c *gin.Context) {
	var (
		err  error
		apiC = response.Api{C: c}
	)
	email := c.Query("email")
	newEmail, err := valueobject.NewEmail(email)
	if email == "" {
		apiC.Response(e.NewError(e.InvalidParam, err))
		return
	}
	if _, err = Dao.GetUser(c, &users.User{Email: newEmail}); err != nil {
		apiC.Response(err)
		return
	}
	codeNum := utils.RandomNum(config.Conf.User.CaptchaLength)
	token, err := Srv.GenCaptchaToken(email, codeNum)
	if err != nil {
		apiC.Response(err)
		return
	}
	if err = Srv.SendCaptchaEmail(email, codeNum); err != nil {
		apiC.Response(err)
		return
	}
	resp := make(map[string]any)
	resp["token"] = token
	resp["expTime"] = config.Conf.User.JwtCaptchaTokenLifeSec
	apiC.Success(resp)
}

func (a Auth) AuthResetPwd(c *gin.Context) {
	var (
		err  error
		apiC = response.Api{C: c}
	)
	email := c.GetString("email")
	if email == "" {
		apiC.Response(e.NewError(e.PermissionDenied, nil))
		return
	}
	codeNum := utils.RandomNum(config.Conf.User.CaptchaLength)
	token, err := Srv.GenCaptchaToken(email, codeNum)
	if err != nil {
		apiC.Response(err)
		return
	}
	if err = Srv.SendCaptchaEmail(email, codeNum); err != nil {
		apiC.Response(err)
		return
	}
	resp := make(map[string]any)
	resp["token"] = token
	resp["expTime"] = config.Conf.User.JwtCaptchaTokenLifeSec
	apiC.Success(resp)
}
