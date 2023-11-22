package auth

import (
	"blog/app/dto"
	"blog/domain/auth/valueobject"
	"blog/infra/e"
	"blog/infra/event"
)

type ServiceAuth struct {
}

func NewServiceAuth() *ServiceAuth {
	return &ServiceAuth{}
}

func (s ServiceAuth) GenCaptchaToken(email, captcha string) (string, error) {
	return valueobject.NewCaptchaToken(email, captcha)
}

func (s ServiceAuth) SendCaptchaEmail(email, captcha string) error {
	//html := mail.NewTemplateForVerifyCode(email, captcha)
	//m := mail.PostMan.NewMail()
	//m.SetHeader("Subject", EmailSubject)
	//m.SetHeader("To", email)
	//m.SetBody("text/html", html.ToString())
	//dialer := mail.PostMan.NewDialer()
	//if err := dialer.DialAndSend(m); err != nil {
	//	return e.NewError(e.EmailSendErr, err)
	//}
	//return nil
	bus.Publish(event.SendMail, event.SendEmailED{
		Email:   email,
		Captcha: captcha,
	})
	return nil
}

func (s ServiceAuth) VaildateCaptcha(cap dto.Captcha) error {
	cc, err := valueobject.ParseCaptchaToken(cap.Token)
	if err != nil {
		return err
	}
	return cc.Verify(cap.Email, cap.Captcha)
}

func (s ServiceAuth) GenAuthToken(user dto.UserDisplay) (string, error) {
	return valueobject.NewAuthToken(user.Uid, user.Email, user.Name, user.Role)
}

func (s ServiceAuth) VaildateAuth(authToken string) (*valueobject.AuthClaims, error) {
	authClaims, err := valueobject.ParseAuthToken(authToken)
	if err != nil {
		return nil, e.NewError(e.TokenVerifyErr, err)
	}
	return authClaims, nil
}
