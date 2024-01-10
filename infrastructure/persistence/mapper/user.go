package mapper

import (
	"blog/domain/aggregate/users"
	"blog/infrastructure/persistence/model"
)

func UserDomainToModel(user users.User) *model.UserM {
	u := user.(*users.UserImpl)
	return &model.UserM{
		Id:       0,
		Uid:      u.Uid,
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		DeleteAt: u.DeleteAt,
	}
}

func UserModelToDomain(user *model.UserM) *users.UserImpl {
	return &users.UserImpl{
		Uid:      user.Uid,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		DeleteAt: user.DeleteAt,
	}
}
