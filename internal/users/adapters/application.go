package adapters

import (
	"blog/pkg/mongodb"
	"users/application"
	"users/application/command"
	"users/application/query"
)

func NewApp() *application.App {
	db := mongodb.GetDB()
	authService := NewUserAuthServce()
	userRepository := NewUserRepository(db)
	return &application.App{
		Commands: application.Commands{
			CreateUser:      command.NewCreateUserHandler(userRepository),
			ResetPassword:   command.NewResetPasswordHandler(userRepository, authService),
			ModifyUserName:  command.NewModifyUserNameHandler(userRepository, authService),
			ModifyUserRoles: command.NewModifyUserRolesHandler(userRepository, authService),
		},
		Queries: application.Queries{
			GetUserInfo:    query.NewGetUserInfoHandler(userRepository),
			GetTokenByUser: query.NewGetUserAuthenticationHandler(userRepository, authService),
		},
	}
}
