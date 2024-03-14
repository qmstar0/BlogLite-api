package application

import (
	"categorys/application/command"
	"categorys/application/query"
)

type App struct {
	Commands Commands
	Queries  Queries
}

type Commands struct {
	CreateCategory command.CreateCategoryHandler
	UpdateCategory command.UpdateCategoryHandler
	DeleteCategory command.DeleteCategoryHandler
}

type Queries struct {
	GetCategory    query.GetCategoryHandler
	GetAllCategory query.GetAllCategoryHandler
}
