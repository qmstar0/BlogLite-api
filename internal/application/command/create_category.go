package command

import (
	"context"
	"go-blog-ddd/internal/application/e"
	"go-blog-ddd/internal/domain/aggregates"
	"go-blog-ddd/internal/domain/commands"
	"go-blog-ddd/internal/domain/values"
)

type CreateCategoryHandler struct {
	repo aggregates.CategoryRepository
}

func NewCreateCategoryHandler(repo aggregates.CategoryRepository) CreateCategoryHandler {
	if repo == nil {
		panic("missing CategoryRepository")
	}
	return CreateCategoryHandler{repo: repo}
}

func (c CreateCategoryHandler) Handle(ctx context.Context, cmd commands.CreateCategory) (uint32, error) {
	name, err := values.NewCategoryName(cmd.Name)
	if err != nil {
		return 0, err
	}

	if find, err := c.repo.FindByName(ctx, name); err != nil {
		return 0, err
	} else if find != nil {
		return 0, e.ResourceAlreadyExists
	}

	nextID := c.repo.NextID(ctx)
	cate := aggregates.NewCategory(nextID, name, cmd.Desc)

	if err = c.repo.Save(ctx, cate); err != nil {
		return 0, err
	}
	return nextID, nil
}
