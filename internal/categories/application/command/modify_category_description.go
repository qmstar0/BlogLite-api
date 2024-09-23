package command

import (
	"context"
	"github.com/qmstar0/BlogLite-api/internal/categories/domain/categories"
	"github.com/qmstar0/BlogLite-api/internal/common/e"
)

type ModifyCategoryDescription struct {
	CategorySlug string
	Description  string
}

type ModifyCategoryDescriptionHandler struct {
	repo categories.CategoryRepository
}

func NewModifyCategoryDescriptionHandler(repo categories.CategoryRepository) *ModifyCategoryDescriptionHandler {
	return &ModifyCategoryDescriptionHandler{repo: repo}
}

func (h *ModifyCategoryDescriptionHandler) Handle(ctx context.Context, cmd ModifyCategoryDescription) error {

	found, err := h.repo.Find(ctx, cmd.CategorySlug)
	if err != nil {
		return err
	} else if found == nil {
		return e.ResourceDoesNotExist
	}

	err = found.ModifyDescription(cmd.Description)
	if err != nil {
		return err
	}

	return h.repo.Save(ctx, found)
}
