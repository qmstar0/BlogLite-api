package command

import (
	"context"
	"github.com/qmstar0/BlogLite-api/internal/categories/domain/categories"
	"github.com/qmstar0/BlogLite-api/internal/common/e"
)

type DeleteCategory struct {
	CategorySlug string
}

type DeleteCategoryHandler struct {
	repo categories.CategoryRepository
}

func NewDeleteCategoryHandler(repo categories.CategoryRepository) *DeleteCategoryHandler {
	return &DeleteCategoryHandler{repo: repo}
}

func (h *DeleteCategoryHandler) Handle(ctx context.Context, cmd DeleteCategory) error {
	found, err := h.repo.Find(ctx, cmd.CategorySlug)
	if err != nil {
		return err
	} else if found == nil {
		return e.ResourceDoesNotExist
	}
	return h.repo.Remove(ctx, found)
}
