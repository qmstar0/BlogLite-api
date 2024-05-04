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

	CreatePost        command.CreatePostHandler
	DeletePost        command.DeletePostHandler
	ModifyPostTags    command.ModifyPostTagsHandler
	ModifyPostVisible command.ModifyPostVisibleHandler
	ModifyPostInfo    command.ModifyPostInfoHandler
	ResetPostCategory command.ResetPostCategoryHandler
	ResetPostContent  command.ResetPostContentHandler
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

			CreatePost:        command.NewCreatePostHandler(postRepo),
			DeletePost:        command.NewDeletePostHandler(postRepo),
			ModifyPostTags:    command.NewModifyPostTagsHandler(postRepo),
			ModifyPostVisible: command.NewModifyPostVisibleHandler(postRepo),
			ModifyPostInfo:    command.NewModifyPostInfoHandler(postRepo),
			ResetPostCategory: command.NewResetPostCategoryHandler(postRepo, categoryService),
			ResetPostContent:  command.NewResetPostContentHandler(postRepo),
		},
		Queries: Queries{
			Posts:     query.NewPostQueryControl(domain.NewPostReadModel(tctx)),
			Categorys: query.NewCategoryQueryControl(domain.NewCategoryReadModel(tctx)),
			Tags:      query.NewTagQueryControl(domain.NewTagReadModel(tctx)),
		},
	}
}
