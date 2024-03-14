package command

import (
	"common/handler"
	"context"
	"users/domain/user"
)

type ModifyUserRoles struct {
	Roles uint16
	Token string
}

type ModifyUserRolesHandler handler.CommandHandler[ModifyUserRoles]

type modifyUserRolesHandler struct {
	userRepo    user.UserRepository
	userAuthSer user.UserAuthService
}

func NewModifyUserRolesHandler(userRepo user.UserRepository, userAuthSer user.UserAuthService) ModifyUserRolesHandler {
	return &modifyUserRolesHandler{userRepo: userRepo, userAuthSer: userAuthSer}
}

func (m modifyUserRolesHandler) Handle(ctx context.Context, cmd ModifyUserRoles) error {
	//TODO implement me
	panic("implement me")
}
