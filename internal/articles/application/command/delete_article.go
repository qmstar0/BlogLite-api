package command

import (
	"context"
	"github.com/qmstar0/BlogLite-api/internal/articles/domain/articles"
	"github.com/qmstar0/BlogLite-api/internal/common/e"
)

type DeleteArticle struct {
	Uri string
}

type DeleteArticleHandler struct {
	resp articles.ArticleRepository
}

func NewDeleteArticleHandler(resp articles.ArticleRepository) *DeleteArticleHandler {
	return &DeleteArticleHandler{resp: resp}
}

func (h DeleteArticleHandler) Handle(ctx context.Context, cmd DeleteArticle) error {
	uri := articles.NewUri(cmd.Uri)
	if err := uri.CheckFormat(); err != nil {
		return err
	}

	if found, err := h.resp.Find(ctx, uri); err != nil {
		return err
	} else if found == nil {
		return e.ResourceDoesNotExist
	} else {
		found.Delete()
		return h.resp.Remove(ctx, found)
	}
}
