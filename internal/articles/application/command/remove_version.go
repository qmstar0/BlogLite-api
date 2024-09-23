package command

import (
	"context"
	"github.com/qmstar0/BlogLite-api/internal/articles/domain/articles"
)

type RemoveVersion struct {
	Uri     string
	Version string
}

type RemoveVersionHandler struct {
	repo articles.ArticleRepository
}

func NewRemoveVersionHandler(repo articles.ArticleRepository) *RemoveVersionHandler {
	return &RemoveVersionHandler{repo: repo}
}

func (h RemoveVersionHandler) Handle(ctx context.Context, cmd RemoveVersion) error {
	uri := articles.NewUri(cmd.Uri)
	if err := uri.CheckFormat(); err != nil {
		return err
	}
	return h.repo.UpdateArticle(ctx, uri, func(article *articles.Article) (*articles.Article, error) {
		err := article.RemoveVersion(cmd.Version)
		if err != nil {
			return nil, err
		}
		return article, nil
	})
}
