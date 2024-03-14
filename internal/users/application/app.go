package application

import (
	"users/application/command"
	"users/application/query"
)

type App struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateUser      command.CreateUserHandler
	ResetPassword   command.ResetPasswordHandler
	ModifyUserName  command.ModifyUserNameHandler
	ModifyUserRoles command.ModifyUserRolesHandler
}

type Queries struct {
	GetUserInfo    query.GetUserInfoHandler
	GetTokenByUser query.GetUserAuthenticationHandler
}
