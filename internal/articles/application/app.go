package application

import (
	"github.com/qmstar0/BlogLite-api/internal/articles/application/command"
	"github.com/qmstar0/BlogLite-api/internal/articles/application/query"
)

type App struct {
	Command Command
	Query   Query
}

type Command struct {
	InitializationArticle   *command.InitializationArticleHandler
	RemoveVersion           *command.RemoveVersionHandler
	ModifyArticleTags       *command.ModifyArticleTagsHandler
	ChangeArticleVisibility *command.ChangeArticleVisibilityHandler
	DeleteArticle           *command.DeleteArticleHandler
	SetArticleVersion       *command.SetArticleVersionHandler
	AddNewVersion           *command.AddNewVersionHandler
	ChangeArticleCategory   *command.ChangeArticleCategoryHandler
}

type Query struct {
	TagList            *query.TagListHandler
	ArticleList        *query.ArticleListHandler
	ArticleDetail      *query.ArticleDetailhandler
	ArticleVersionList *query.ArticleVersionListHandler
}
