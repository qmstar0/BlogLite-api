package command

import (
	"context"
	"go-blog-ddd/internal/adapter/e"
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
		return 0, e.DErrInvalidOperation.WithError(err)
	}

	//if find, err := c.repo.FindByName(ctx, name); err != nil {
	//	return 0, err
	//} else if find != nil {
	//	return 0, e.RErrResourceExists
	//}

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
