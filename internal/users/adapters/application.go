package adapters

import (
	"users/application"
)

func NewApp() *application.App {

	return &application.App{
		Commands: application.Commands{
			UpdateUsername: nil,
			ResetPassword:  nil,
		},
		Queries: application.Queries{},
	}
}
