package command

import (
	"context"
	"go-blog-ddd/internal/domain/aggregates"
	"go-blog-ddd/internal/domain/commands"
)

type DeleteCategoryHandler struct {
	cateRepo aggregates.CategoryRepository
}

func NewDeleteCategoryHandler(repo aggregates.CategoryRepository) DeleteCategoryHandler {
	if repo == nil {
		panic("missing CategoryRepository")
	}
	return DeleteCategoryHandler{cateRepo: repo}
}

func (u DeleteCategoryHandler) Handle(ctx context.Context, cmd commands.DeleteCategory) error {
	return u.cateRepo.Delete(ctx, cmd.ID)
}
