package apps

import (
	"context"
	"github.com/qmstar0/nightsky-api/internal/apps/commandhandler"
	"github.com/qmstar0/nightsky-api/internal/apps/query"
	"github.com/qmstar0/nightsky-api/internal/domain/services"
	"github.com/qmstar0/nightsky-api/internal/pkg/domain"
	"github.com/qmstar0/nightsky-api/pkg/postgresql"
	"github.com/qmstar0/nightsky-api/pkg/transaction"
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

func NewDomainApp() *DomainApp {
	tctx := transaction.NewTransactionContext(postgresql.GetDB())
	categoryRepo := domain.NewCategoryRepository(tctx)
	categoryService := services.NewPostDomainService(categoryRepo)
	postRepo := domain.NewPostRepository(tctx)
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
			Posts:     domain.NewPostReadModel(tctx),
			Categorys: domain.NewCategoryReadModel(tctx),
			Tags:      domain.NewTagReadModel(tctx),
		},
	}
}
