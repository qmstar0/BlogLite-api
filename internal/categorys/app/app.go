package app

import "blog/internal/categorys/app/command"

type Application struct {
}

type Commands struct {
	CreateCategory command.CreateCategoryHandler
}

type Queries struct {
}
