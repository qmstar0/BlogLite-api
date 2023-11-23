package handlers

import (
	"blog/app/dto"
	"blog/app/response"
	"blog/domain/users"
	"blog/domain/users/valueobject"
	"blog/infra/config"
	"blog/infra/e"
	"blog/router"
	"github.com/gin-gonic/gin"
)

type UserHandler struct {
}

func NewUser() router.User {
	return &UserHandler{}
}

func (u UserHandler) Login(c *gin.Context) {
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
	if err := Srv.VaildateCaptcha(Captcha); err != nil {
		apiC.Response(err)
		return
	}
	email, err := valueobject.NewEmail(Captcha.Email)
	if err != nil {
		apiC.Response(err)
		return
	}
	user, err := Dao.GetUser(c, &users.User{Email: email})
	if err != nil {
		apiC.Response(err)
		return
	}
	token, err := user.GenUserAuthToken()
	if err != nil {
		apiC.Response(err)
		return
	}
	resp := make(map[string]any)
	resp["token"] = token
	resp["expTime"] = config.Conf.User.JwtAuthTokenLifeDay
	apiC.Success(resp)
}

func (u UserHandler) Register(c *gin.Context) {
	var (
		err  error
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
	if err := Srv.VaildateCaptcha(Captcha); err != nil {
		apiC.Response(err)
		return
	}
	newUser, err := users.NewUser(Captcha.Email, 1)
	if err != nil {
		apiC.Response(err)
		return
	}
	err = Dao.NewUser(c, newUser)
	if err != nil {
		apiC.Response(err)
		return
	}
	token, err := newUser.GenUserAuthToken()
	if err != nil {
		apiC.Response(err)
		return
	}
	resp := make(map[string]any)
	resp["token"] = token
	resp["expTime"] = config.Conf.User.JwtAuthTokenLifeDay
	apiC.Success(resp)
}

func (u UserHandler) Update(c *gin.Context) {
	var (
		apiC = response.Api{C: c}
	)
	emailStr := c.GetString("email")
	if emailStr == "" {
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
	email, err := valueobject.NewEmail(emailStr)
	if err != nil {
		apiC.Response(err)
		return
	}
	if err := Dao.UptUser(c, &users.User{
		Email:    email,
		UserName: userStore.UserName,
	}); err != nil {
		apiC.Response(err)
		return
	}
	apiC.Success(nil)
}

func (u UserHandler) ResetPwd(c *gin.Context) {
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
	if err := Srv.VaildateCaptcha(Captcha); err != nil {
		apiC.Response(err)
		return
	}
	newPwd := c.PostForm("newPwd")
	if newPwd == "" {
		apiC.Response(e.NewError(e.MissingParam, nil))
		return
	}
	email, err := valueobject.NewEmail(Captcha.Email)
	if err != nil {
		apiC.Response(err)
		return
	}
	user := &users.User{Email: email}
	if err = user.ResetPassword(newPwd); err != nil {
		apiC.Response(err)
		return
	}
	if err := Dao.UptUser(c, user); err != nil {
		apiC.Response(err)
		return
	}
	apiC.Success(nil)
}
