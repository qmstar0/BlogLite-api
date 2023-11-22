package handlers

import (
	"blog/app/dto"
	"blog/app/response"
	"blog/app/service"
	"blog/infra/config"
	"blog/infra/e"
	"blog/router"
	"blog/utils"
	"github.com/gin-gonic/gin"
)

type userDTO interface {
	dto.UserR
	dto.UserW
	dto.Authorizer
}

type User struct {
	Srv userDTO
}

func NewUser() router.User {
	return &User{Srv: service.GetSrv()}
}

func NewAuth() router.Auth {
	return &User{Srv: service.GetSrv()}
}

func (u User) AuthRegister(c *gin.Context) {
	var (
		err  error
		apiC = response.Api{C: c}
	)
	email := c.Query("email")
	if email == "" {
		apiC.Response(e.NewError(e.InvalidParam, nil))
		return
	}
	if _, err = u.Srv.GetUser(c, email); err == nil {
		apiC.Response(e.NewError(e.UserDuplicateCreationErr, nil))
		return
	}
	codeNum := utils.RandomNum(config.Conf.User.CaptchaLength)
	token, err := u.Srv.GenCaptchaToken(email, codeNum)
	if err != nil {
		apiC.Response(err)
		return
	}
	if err = u.Srv.SendCaptchaEmail(email, codeNum); err != nil {
		apiC.Response(err)
		return
	}
	resp := make(map[string]any)
	resp["token"] = token
	resp["expTime"] = config.Conf.User.JwtCaptchaTokenLifeSec
	apiC.Success(resp)
}

func (u User) AuthLogin(c *gin.Context) {
	var (
		err  error
		apiC = response.Api{C: c}
	)
	email := c.Query("email")
	if email == "" {
		apiC.Response(e.NewError(e.InvalidParam, nil))
		return
	}
	if _, err = u.Srv.GetUser(c, email); err != nil {
		apiC.Response(err)
		return
	}
	codeNum := utils.RandomNum(config.Conf.User.CaptchaLength)
	token, err := u.Srv.GenCaptchaToken(email, codeNum)
	if err != nil {
		apiC.Response(err)
		return
	}
	if err = u.Srv.SendCaptchaEmail(email, codeNum); err != nil {
		apiC.Response(err)
		return
	}
	resp := make(map[string]any)
	resp["token"] = token
	resp["expTime"] = config.Conf.User.JwtCaptchaTokenLifeSec
	apiC.Success(resp)
}

func (u User) AuthResetPwd(c *gin.Context) {
	var (
		err  error
		apiC = response.Api{C: c}
	)
	email := c.GetString("email")
	if email == "" {
		apiC.Response(e.NewError(e.InvalidParam, nil))
		return
	}
	codeNum := utils.RandomNum(config.Conf.User.CaptchaLength)
	token, err := u.Srv.GenCaptchaToken(email, codeNum)
	if err != nil {
		apiC.Response(err)
		return
	}
	if err = u.Srv.SendCaptchaEmail(email, codeNum); err != nil {
		apiC.Response(err)
		return
	}
	resp := make(map[string]any)
	resp["token"] = token
	resp["expTime"] = config.Conf.User.JwtCaptchaTokenLifeSec
	apiC.Success(resp)
}

func (u User) Login(c *gin.Context) {
	var (
		apiC = response.Api{C: c}
	)
	auth, exists := c.Get("captcha")
	if !exists {
		apiC.Response(e.NewError(e.InvalidParam, nil))
		return
	}
	Captcha, ok := auth.(dto.Captcha)
	if !ok {
		apiC.Response(e.NewError(e.InvalidParam, nil))
		return
	}
	if err := u.Srv.VaildateCaptcha(Captcha); err != nil {
		apiC.Response(err)
		return
	}
	user, err := u.Srv.GetUser(c, Captcha.Email)
	if err != nil {
		apiC.Response(err)
		return
	}
	token, err := u.Srv.GenAuthToken(user)
	if err != nil {
		apiC.Response(err)
		return
	}
	resp := make(map[string]any)
	resp["token"] = token
	resp["expTime"] = config.Conf.User.JwtAuthTokenLifeDay
	apiC.Success(resp)
}

func (u User) Register(c *gin.Context) {
	var (
		apiC = response.Api{C: c}
	)
	auth, exists := c.Get("captcha")
	if !exists {
		apiC.Response(e.NewError(e.InvalidParam, nil))
		return
	}
	authDTO, ok := auth.(dto.Captcha)
	if !ok {
		apiC.Response(e.NewError(e.InvalidParam, nil))
		return
	}
	if err := u.Srv.VaildateCaptcha(authDTO); err != nil {
		apiC.Response(err)
		return
	}
	user, err := u.Srv.NewUser(c, authDTO.Email)
	if err != nil {
		apiC.Response(err)
		return
	}
	token, err := u.Srv.GenAuthToken(user)
	if err != nil {
		apiC.Response(err)
		return
	}
	resp := make(map[string]any)
	resp["token"] = token
	resp["expTime"] = config.Conf.User.JwtAuthTokenLifeDay
	apiC.Success(resp)
}

func (u User) Update(c *gin.Context) {
	var (
		apiC = response.Api{C: c}
	)
	email := c.GetString("email")
	if email == "" {
		apiC.Response(e.NewError(e.InvalidParam, nil))
		return
	}
	store, exists := c.Get("store")
	if !exists {
		apiC.Response(e.NewError(e.InvalidParam, nil))
		return
	}
	userStore, ok := store.(dto.UserStore)
	if !ok {
		apiC.Response(e.NewError(e.InvalidParam, nil))
		return
	}
	if err := u.Srv.UpdateUser(c, email, userStore); err != nil {
		apiC.Response(err)
		return
	}
	apiC.Success(nil)
}

func (u User) ResetPwd(c *gin.Context) {
	var (
		apiC = response.Api{C: c}
	)
	auth, exists := c.Get("captcha")
	if !exists {
		apiC.Response(e.NewError(e.InvalidParam, nil))
		return
	}
	captcha, ok := auth.(dto.Captcha)
	if !ok {
		apiC.Response(e.NewError(e.InvalidParam, nil))
		return
	}
	if err := u.Srv.VaildateCaptcha(captcha); err != nil {
		apiC.Response(err)
		return
	}
	newPwd := c.PostForm("newPwd")
	if newPwd == "" {
		apiC.Response(e.NewError(e.MissingParam, nil))
		return
	}
	if err := u.Srv.ResetPwd(c, captcha.Email, newPwd); err != nil {
		apiC.Response(err)
		return
	}
	apiC.Success(nil)
}
