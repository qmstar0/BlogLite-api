package command

import (
	"context"
	"github.com/qmstar0/BlogLite-api/internal/articles/domain/articles"
)

type SetArticleVersion struct {
	Uri     string
	Version string
}

type SetArticleVersionHandler struct {
	repo articles.ArticleRepository
}

func NewSetArticleVersionHandler(repo articles.ArticleRepository) *SetArticleVersionHandler {
	return &SetArticleVersionHandler{repo: repo}
}

func (h SetArticleVersionHandler) Handle(ctx context.Context, cmd SetArticleVersion) error {
	uri := articles.NewUri(cmd.Uri)
	if err := uri.CheckFormat(); err != nil {
		return err
	}

	return h.repo.UpdateArticle(ctx, uri, func(article *articles.Article) (*articles.Article, error) {
		err := article.SetCurrentVersion(cmd.Version)
		if err != nil {
			return nil, err
		}
		return article, nil
	})
}
