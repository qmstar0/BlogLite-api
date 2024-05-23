package apps

import (
	"context"
	"go-blog-ddd/internal/apps/commandhandler"
	"go-blog-ddd/internal/apps/query"
	"go-blog-ddd/internal/domain/services"
	domain2 "go-blog-ddd/internal/pkg/domain"
	"go-blog-ddd/pkg/postgresql"
	"go-blog-ddd/pkg/transaction"
)

type DomainApp struct {
	Commands Commands
	Queries  Queries

	transaction transaction.TransactionContext
}

func (a DomainApp) Transaction(ctx context.Context, fn func(tctx context.Context) error) error {
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
	CreateCategory     commandhandler.CreateCategoryHandler
	ModifyCategoryDesc commandhandler.ModifyCategoryDescHandler
	DeleteCategory     commandhandler.DeleteCategoryHandler

	CreatePost       commandhandler.CreatePostHandler
	DeletePost       commandhandler.DeletePostHandler
	ModifyPost       commandhandler.ModifyPostHandler
	ResetPostContent commandhandler.ResetPostContentHandler
}

func NewDomainControl() *DomainApp {
	tctx := transaction.NewTransactionContext(postgresql.GetDB())
	categoryRepo := domain2.NewCategoryRepository(tctx)
	categoryService := services.NewPostDomainService(categoryRepo)
	postRepo := domain2.NewPostRepository(tctx)
	return &DomainApp{
		transaction: tctx,
		Commands: Commands{
			CreateCategory:     commandhandler.NewCreateCategoryHandler(categoryRepo),
			ModifyCategoryDesc: commandhandler.NewModifyCategoryDescHandler(categoryRepo),
			DeleteCategory:     commandhandler.NewDeleteCategoryHandler(categoryRepo),

			CreatePost:       commandhandler.NewCreatePostHandler(postRepo),
			DeletePost:       commandhandler.NewDeletePostHandler(postRepo),
			ModifyPost:       commandhandler.NewModifyPostHandler(postRepo, categoryService),
			ResetPostContent: commandhandler.NewResetPostContentHandler(postRepo),
		},
		Queries: Queries{
			Posts:     domain2.NewPostReadModel(tctx),
			Categorys: domain2.NewCategoryReadModel(tctx),
			Tags:      domain2.NewTagReadModel(tctx),
		},
	}
}
