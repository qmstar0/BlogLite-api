package application

import (
	"context"
	"go-blog-ddd/internal/adapter/domain"
	"go-blog-ddd/internal/adapter/postgresql"
	"go-blog-ddd/internal/adapter/transaction"
	"go-blog-ddd/internal/application/command"
	"go-blog-ddd/internal/application/query"
	"go-blog-ddd/internal/domain/services"
)

type App struct {
	Commands Commands
	Queries  Queries

	transaction transaction.TransactionContext
}

func (a App) Transaction(ctx context.Context, fn func(tctx context.Context) error) error {
	return a.transaction.Transaction(ctx, func(tctx transaction.TransactionContext) error {
		return fn(tctx)
	})
}

type Queries struct {
	Posts     query.PostQueryControl
	Categorys query.CategoryQueryControl
	Tags      query.TagQueryControl
}

type Commands struct {
	CreateCategory     command.CreateCategoryHandler
	ModifyCategoryDesc command.ModifyCategoryDescHandler
	DeleteCategory     command.DeleteCategoryHandler

	CreatePost       command.CreatePostHandler
	DeletePost       command.DeletePostHandler
	ModifyPost       command.ModifyPostHandler
	ResetPostContent command.ResetPostContentHandler
}

func NewApp() *App {
	tctx := transaction.NewTransactionContext(postgresql.GetDB())
	categoryRepo := domain.NewCategoryRepository(tctx)
	categoryService := services.NewPostDomainService(categoryRepo)
	postRepo := domain.NewPostRepository(tctx)
	return &App{
		transaction: tctx,
		Commands: Commands{
			CreateCategory:     command.NewCreateCategoryHandler(categoryRepo),
			ModifyCategoryDesc: command.NewModifyCategoryDescHandler(categoryRepo),
			DeleteCategory:     command.NewDeleteCategoryHandler(categoryRepo),

			CreatePost:       command.NewCreatePostHandler(postRepo),
			DeletePost:       command.NewDeletePostHandler(postRepo),
			ModifyPost:       command.NewModifyPostHandler(postRepo, categoryService),
			ResetPostContent: command.NewResetPostContentHandler(postRepo),
		},
		Queries: Queries{
			Posts:     domain.NewPostReadModel(tctx),
			Categorys: domain.NewCategoryReadModel(tctx),
			Tags:      domain.NewTagReadModel(tctx),
		},
	}
}
