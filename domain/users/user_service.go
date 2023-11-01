package users

import (
	"blog/app/dto"
	"blog/domain/users/token"
	"blog/domain/users/valueobject"
	"blog/infra/e"
	"blog/infra/mail"
	"context"
)

const EmailSubject = "Please Confirm Your email"

// ServiceUser 用户领域服务
type ServiceUser struct {
	repo RepoUser
}

func NewServiceUser(repo RepoUser) *ServiceUser {
	return &ServiceUser{repo: repo}
}

func (s ServiceUser) GetUserByUid(c context.Context, uid string) (dto.UserDisplay, error) {
	user, err := s.repo.GetUser(c, &User{Uid: uid})
	if err != nil {
		return dto.UserDisplay{}, err
	}
	return dto.UserDisplay{
		Uid:   user.Uid,
		Name:  user.UserName,
		Email: user.Email.ToString(),
		Role:  user.Role.ToUint(),
	}, nil
}

func (s ServiceUser) GenCaptchaToken(email, captcha string) (string, error) {
	return token.NewCaptchaToken(email, captcha)
}

func (s ServiceUser) SendCaptchaEmail(email, captcha string) error {
	html := mail.NewTemplateForVerifyCode(email, captcha)
	m := mail.PostMan.NewMail()
	m.SetHeader("Subject", EmailSubject)
	m.SetHeader("To", email)
	m.SetBody("text/html", html.ToString())
	dialer := mail.PostMan.NewDialer()
	if err := dialer.DialAndSend(m); err != nil {
		return e.NewError(e.EmailSendErr, err)
	}
	return nil
}

func (s ServiceUser) VaildateCaptcha(cap dto.Captcha) error {
	cc, err := token.ParseCaptchaToken(cap.Token)
	if err != nil {
		return err
	}
	return cc.Verify(cap.Email, cap.Captcha)
}

func (s ServiceUser) GenAuthToken(user dto.UserDisplay) (string, error) {
	return token.NewAuthToken(user.Uid, user.Email, user.Name, user.Role)
}

func (s ServiceUser) VaildateAuth(authToken string) (*token.AuthClaims, error) {
	authClaims, err := token.ParseAuthToken(authToken)
	if err != nil {
		return nil, e.NewError(e.TokenVerifyErr, err)
	}
	return authClaims, nil
}

func (s ServiceUser) GetUser(c context.Context, email string) (dto.UserDisplay, error) {
	newEmail, err := valueobject.NewEmail(email)
	if err != nil {
		return dto.UserDisplay{}, err
	}
	user, err := s.repo.GetUser(c, &User{Email: newEmail})
	if err != nil {
		return dto.UserDisplay{}, err
	}
	return dto.UserDisplay{
		Uid:   user.Uid,
		Name:  user.UserName,
		Email: user.Email.ToString(),
		Role:  user.Role.ToUint(),
	}, nil
}

func (s ServiceUser) NewUser(c context.Context, email string) (dto.UserDisplay, error) {
	u, err := NewUser(email, valueobject.CommonUserTag)
	if err != nil {
		return dto.UserDisplay{}, err
	}
	user, err := s.repo.GetUser(c, u)
	if err != nil {
		return dto.UserDisplay{}, err
	}
	return dto.UserDisplay{
		Uid:   user.Uid,
		Name:  user.UserName,
		Email: user.Email.ToString(),
		Role:  user.Role.ToUint(),
	}, nil
}

func (s ServiceUser) UpdateUser(c context.Context, email string, store dto.UserStore) error {
	newEmail, err := valueobject.NewEmail(email)
	if err != nil {
		return err
	}
	user, err := s.repo.GetUser(c, &User{Email: newEmail})
	if err = user.UpdateUserName(store.UserName); err != nil {
		return err
	}
	if err = s.repo.UptUser(c, user); err != nil {
		return err
	}
	return nil
}

func (s ServiceUser) ResetPwd(c context.Context, email, newPwd string) error {
	newEmail, err := valueobject.NewEmail(email)
	if err != nil {
		return err
	}
	user, err := s.repo.GetUser(c, &User{Email: newEmail})
	if err = user.ResetPassword(newPwd); err != nil {
		return err
	}
	if err = s.repo.UptUser(c, user); err != nil {
		return err
	}
	return nil
}
