package command

import (
	"context"
	"github.com/qmstar0/BlogLite-api/internal/articles/domain/articles"
	"github.com/qmstar0/BlogLite-api/internal/common/e"
)

type InitializationArticle struct {
	Uri        string
	CategoryID string
}

type InitializationArticleHandler struct {
	resp articles.ArticleRepository
	ser  CategoryValidityCheckService
}

func NewInitializationArticleHandler(resp articles.ArticleRepository, ser CategoryValidityCheckService) *InitializationArticleHandler {
	return &InitializationArticleHandler{resp: resp, ser: ser}
}

func (h InitializationArticleHandler) Handle(ctx context.Context, cmd InitializationArticle) error {
	uri := articles.NewUri(cmd.Uri)
	if err := uri.CheckFormat(); err != nil {
		return err
	}

	if err := h.ser.CategoryExist(ctx, cmd.CategoryID); err != nil {
		return err
	}

	if found, err := h.resp.Find(ctx, uri); err != nil {
		return err
	} else if found != nil {
		return e.InvalidActionError("该文章资源已存在，请勿重复初始化")
	}

	return h.resp.Save(ctx, articles.NewArticle(uri, cmd.CategoryID))
}
