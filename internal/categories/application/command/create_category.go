package command

import (
	"context"
	"github.com/qmstar0/BlogLite-api/internal/categories/domain/categories"
	"github.com/qmstar0/BlogLite-api/internal/common/e"
)

type CreateCategory struct {
	Slug        string
	Name        string
	Description string
}

type CreateCategoryHandler struct {
	repo categories.CategoryRepository
}

func NewCreateCategoryHandler(repo categories.CategoryRepository) *CreateCategoryHandler {
	return &CreateCategoryHandler{repo: repo}
}

func (h *CreateCategoryHandler) Handle(ctx context.Context, cmd CreateCategory) error {
	found, err := h.repo.Find(ctx, cmd.Slug)
	if err != nil {
		return err
	}
	if found != nil {
		return e.InvalidActionError("资源已存在")
	}

	exist, err := h.repo.CheckNameExist(ctx, cmd.Name)
	if err != nil {
		return err
	}

	if exist {
		return e.InvalidActionError("该分类命名已存在，请使用别的命名")
	}

	return h.repo.Save(ctx, categories.NewCategory(cmd.Slug, cmd.Name, cmd.Description))
}
