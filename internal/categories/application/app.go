package application

import (
	"github.com/qmstar0/BlogLite-api/internal/categories/application/command"
	"github.com/qmstar0/BlogLite-api/internal/categories/application/query"
)

type App struct {
	Command Command
	Query   Query
}

type Command struct {
	CreateCategory            *command.CreateCategoryHandler
	ModifyCategoryDescription *command.ModifyCategoryDescriptionHandler
	DeleteCategory            *command.CheckAndDeleteCategoryHandler
}

type Query struct {
	CategoryList *query.CategoryListHandler
}
