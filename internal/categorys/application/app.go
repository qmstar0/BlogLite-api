package application

import (
	"categorys/application/command"
	"categorys/application/event"
)

type App struct {
	Commands Commands
	Queries  Queries
	Events   Events
}

type Commands struct {
	CreateCategory command.CreateCategoryHandler
}

type Queries struct {
}

type Events struct {
	CategoryCreated event.CategoryCreatedHandler
}
