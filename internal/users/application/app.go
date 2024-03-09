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
	UpdateUsername command.UpdateUsernameHandler
	ResetPassword  command.ResetPasswordHandler
}

type Queries struct {
	GetUser query.GetUserHandler
}
