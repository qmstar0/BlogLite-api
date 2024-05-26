package commandhandler

import (
	"context"
	"github.com/qmstar0/domain/internal/domain/aggregates"
	"github.com/qmstar0/domain/internal/domain/commands"
	"github.com/qmstar0/domain/internal/domain/values"
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
	if err = c.repo.ResourceUniquenessCheck(ctx, name); err != nil {
		return 0, err
	}

	nextID, err := c.repo.NextID(ctx)
	if err != nil {
		return 0, err
	}
	cate := aggregates.NewCategory(nextID, name, cmd.Desc)
	err = c.repo.Save(ctx, cate)
	if err != nil {
		return 0, err
	}
	return nextID, nil
}
