package command

import (
	"context"
	"github.com/qmstar0/BlogLite-api/internal/articles/domain/articles"
)

type ModifyArticleTags struct {
	Uri  string
	Tags []string
}

type ModifyArticleTagsHandler struct {
	repo articles.ArticleRepository
}

func NewModifyArticleTagsHandler(repo articles.ArticleRepository) *ModifyArticleTagsHandler {
	return &ModifyArticleTagsHandler{repo: repo}
}

func (h ModifyArticleTagsHandler) Handle(ctx context.Context, cmd ModifyArticleTags) error {
	uri := articles.NewUri(cmd.Uri)
	if err := uri.CheckFormat(); err != nil {
		return err
	}

	group, err := articles.NewTagGroup(cmd.Tags)
	if err != nil {
		return err
	}

	return h.repo.UpdateArticle(ctx, uri, func(article *articles.Article) (*articles.Article, error) {
		article.ChangeTagGroup(group)
		return article, nil
	})
}
