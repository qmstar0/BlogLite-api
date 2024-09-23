package service

import (
	"context"
	"github.com/qmstar0/BlogLite-api/internal/categories/adapter"
	"github.com/qmstar0/BlogLite-api/internal/categories/application"
	"github.com/qmstar0/BlogLite-api/internal/categories/application/command"
	"github.com/qmstar0/BlogLite-api/internal/categories/application/query"
	"github.com/qmstar0/BlogLite-api/pkg/postgresql"
)

func NewComponentTestApplication(ctx context.Context) *application.App {
	return newApplication(ctx, MockGetCategoryUsedService{})
}

func NewApplication(ctx context.Context) *application.App {
	db := postgresql.GetDB()

	service := adapter.NewGetCategoryUsedService(db)
	return newApplication(ctx, service)
}

func newApplication(ctx context.Context, service command.GetCategoryUsedService) *application.App {

	db := postgresql.GetDB()

	repo := adapter.NewPostgresCategoryRepository(db)

	return &application.App{
		Command: application.Command{
			CreateCategory:            command.NewCreateCategoryHandler(repo),
			ModifyCategoryDescription: command.NewModifyCategoryDescriptionHandler(repo),
			DeleteCategory:            command.NewCheckAndDeleteCategoryHandler(service, command.NewDeleteCategoryHandler(repo)),
		},
		Query: application.Query{
			CategoryList: query.NewCategoryListHandler(repo),
		},
	}
}
