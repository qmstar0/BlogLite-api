package command

import (
	"context"
	"github.com/qmstar0/BlogLite-api/internal/articles/domain/articles"
)

type ChangeArticleCategory struct {
	Uri        string
	CategoryID string
}

type ChangeArticleCategoryHandler struct {
	repo articles.ArticleRepository
	ser  CategoryValidityCheckService
}

func NewChangeArticleCategoryHandler(repo articles.ArticleRepository, ser CategoryValidityCheckService) *ChangeArticleCategoryHandler {
	return &ChangeArticleCategoryHandler{repo: repo, ser: ser}
}

func (h ChangeArticleCategoryHandler) Handle(ctx context.Context, cmd ChangeArticleCategory) error {
	uri := articles.NewUri(cmd.Uri)
	if err := uri.CheckFormat(); err != nil {
		return err
	}

	return h.repo.UpdateArticle(ctx, uri, func(article *articles.Article) (*articles.Article, error) {
		if err := h.ser.CategoryExist(ctx, cmd.CategoryID); err != nil {
			return nil, err
		}
		article.ChangeCategory(cmd.CategoryID)
		return article, nil
	})
}
