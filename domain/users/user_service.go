package users

import (
	"blog/app/dto"
	"blog/domain/users/valueobject"
	"context"
)

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
	if err = s.repo.NewUser(c, u); err != nil {
		return dto.UserDisplay{}, err
	}
	return dto.UserDisplay{
		Uid:   u.Uid,
		Name:  u.UserName,
		Email: u.Email.ToString(),
		Role:  u.Role.ToUint(),
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
