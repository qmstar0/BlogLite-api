package application

import (
	"blog/internal/categorys/application/command"
	"blog/internal/common/bus"
)

type App struct {
	CommandsBus bus.CommandBus
	QueriesBus  bus.QueryBus

	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateCategory command.CreateCategoryHandler
}

type Queries struct {
}
