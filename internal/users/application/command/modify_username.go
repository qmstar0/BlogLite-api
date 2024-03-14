package command

import (
	"common/handler"
	"context"
	"users/domain/user"
)

type ModifyUserName struct {
	Token string
	Name  string
}

type ModifyUserNameHandler handler.CommandHandler[ModifyUserName]

type modifyUserNameHandler struct {
	userRepo    user.UserRepository
	userAuthSer user.UserAuthService
}

func NewModifyUserNameHandler(userRepo user.UserRepository, userAuthSer user.UserAuthService) ModifyUserNameHandler {
	return &modifyUserNameHandler{userRepo: userRepo, userAuthSer: userAuthSer}
}

func (m modifyUserNameHandler) Handle(ctx context.Context, cmd ModifyUserName) error {
	//TODO implement me
	panic("implement me")
}
