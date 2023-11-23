package users

import (
	"blog/app/dto"
	"blog/infra/e"
	"blog/infra/event"
	"blog/infra/jwt"
	"blog/utils"
)

// ServiceUser 用户领域服务
type ServiceUser struct {
}

func NewServiceUser() *ServiceUser {
	return &ServiceUser{}
}
func (s ServiceUser) SendCaptchaEmail(email, captcha string) error {
	bus.Publish(event.SendMail, event.SendEmailED{
		Email:   email,
		Captcha: captcha,
	})
	return nil
}

func (s ServiceUser) VaildateAuth(authToken string) (jwt.MapClaims, error) {
	return jwt.ParseToken(authToken)
}
func (s ServiceUser) GenCaptchaToken(email, captcha string) (string, error) {
	saltHashCaptcha, err := utils.GetSaltHash(captcha)
	if err != nil {
		return "", err
	}
	data := map[string]any{
		"email":       email,
		"hashCaptcha": saltHashCaptcha,
	}
	signStr, err := jwt.Sign(data, userAuthTokenLifeTime)
	if err != nil {
		return "", e.NewError(e.JwtSignErr, err)
	}
	return signStr, nil
}
func (s ServiceUser) VaildateCaptcha(cap dto.Captcha) error {
	mapClaims, err := jwt.ParseToken(cap.Token)
	if err != nil {
		return e.NewError(e.JwtParseErr, err)
	}
	saltHashCaptcha, err := utils.GetSaltHash(cap.Email)

	if mapClaims["email"] != cap.Email && mapClaims["hashCaptcha"] != saltHashCaptcha {
		return e.NewError(e.TokenVerifyErr, err)
	}
	return nil
}
