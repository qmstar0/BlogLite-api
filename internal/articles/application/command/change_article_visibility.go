package command

import (
	"context"
	"github.com/qmstar0/BlogLite-api/internal/articles/domain/articles"
)

type ChangeArticleVisibility struct {
	Uri        string
	Visibility bool
}

type ChangeArticleVisibilityHandler struct {
	resp articles.ArticleRepository
}

func NewChangeArticleVisibilityHandler(resp articles.ArticleRepository) *ChangeArticleVisibilityHandler {
	return &ChangeArticleVisibilityHandler{resp: resp}
}

func (h ChangeArticleVisibilityHandler) Handle(ctx context.Context, cmd ChangeArticleVisibility) error {
	uri := articles.NewUri(cmd.Uri)
	if err := uri.CheckFormat(); err != nil {
		return err
	}
	return h.resp.UpdateArticle(ctx, uri, func(article *articles.Article) (*articles.Article, error) {
		article.ChangeVisibility(cmd.Visibility)
		return article, nil
	})
}
