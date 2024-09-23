package service

import (
	"context"
	"github.com/qmstar0/BlogLite-api/internal/articles/adapter"
	"github.com/qmstar0/BlogLite-api/internal/articles/application"
	"github.com/qmstar0/BlogLite-api/internal/articles/application/command"
	"github.com/qmstar0/BlogLite-api/internal/articles/application/query"
	categoryAdapter "github.com/qmstar0/BlogLite-api/internal/categories/adapter"
	"github.com/qmstar0/BlogLite-api/pkg/postgresql"
	"github.com/qmstar0/shutdown"
)

func NewComponentTestApplication(ctx context.Context) *application.App {
	return newApplication(
		ctx,
		MockCategoryValidityCheckService{},
		adapter.NewMarkdownParser(),
	)
}

func NewApplication(ctx context.Context) *application.App {
	db := postgresql.GetDB()
	categoryValidityCheckService := adapter.NewCategoryValidityCheckService(categoryAdapter.NewPostgresCategoryRepository(db))
	return newApplication(
		ctx,
		categoryValidityCheckService,
		adapter.NewMarkdownParser(),
	)
}

func newApplication(ctx context.Context, categoryService command.CategoryValidityCheckService, markdownService command.MarkdownParseService) *application.App {
	bus := adapter.NewBus()
	//bus.Register("test", mockHandler{})
	defer func() {
		bus.Run(ctx)
		shutdown.RegisterTasks(bus.RouterClose)
	}()

	db := postgresql.GetDB()

	repo := adapter.NewPostgresArticleRepository(db, bus)

	articleDetailReadmodel := adapter.NewPostgresArticleDetailReadmodel(db)
	articleTagReadmodel := adapter.NewPostgresArticleTagReadmodel(db)
	articleVersionReadmodel := adapter.NewPostgresArticleVersionReadmodel(db)

	bus.Register("article-detail-readmodel", articleDetailReadmodel)
	bus.Register("article-tag-readmodel", articleTagReadmodel)
	bus.Register("article-version-readmodel", articleVersionReadmodel)

	return &application.App{
		Command: application.Command{
			InitializationArticle:   command.NewInitializationArticleHandler(repo, categoryService),
			RemoveVersion:           command.NewRemoveVersionHandler(repo),
			ModifyArticleTags:       command.NewModifyArticleTagsHandler(repo),
			ChangeArticleVisibility: command.NewChangeArticleVisibilityHandler(repo),
			DeleteArticle:           command.NewDeleteArticleHandler(repo),
			SetArticleVersion:       command.NewSetArticleVersionHandler(repo),
			AddNewVersion:           command.NewAddNewVersionHandler(repo, markdownService, adapter.NewArticleVersionDuplicationCheckService(db)),
			ChangeArticleCategory:   command.NewChangeArticleCategoryHandler(repo, categoryService),
		},
		Query: application.Query{
			TagList:            query.NewTagListHandler(articleTagReadmodel),
			ArticleList:        query.NewArticleListHandler(articleDetailReadmodel),
			ArticleDetail:      query.NewArticleDetailhandler(articleDetailReadmodel),
			ArticleVersionList: query.NewArticleVersionListHandler(articleVersionReadmodel),
		},
	}
}
