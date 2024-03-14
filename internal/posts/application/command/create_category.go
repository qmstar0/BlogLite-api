package command

import (
	"categorys/domain/v1/category"
	"common/e"
	"common/handler"
	"context"
)

type CreateCategory struct {
	Name    string
	SeoDesc string
}

type CreateCategoryHandler handler.CommandHandler[CreateCategory]

type createCategoryHandler struct {
	cateRepo category.CategoryRepository
}

func NewCreateCategoryHandler(repo category.CategoryRepository) CreateCategoryHandler {
	return &createCategoryHandler{
		cateRepo: repo,
	}
}

func (c createCategoryHandler) Handle(ctx context.Context, cmd CreateCategory) error {
	var err error
	newName, err := category.NewName(cmd.Name)
	if err != nil {
		return e.Wrap(e.NewValueObjectErr, err)
	}

	aggid := newName.ToID()

	if exist, err := c.cateRepo.Exist(ctx, aggid); err != nil {
		return e.Wrap(e.FindEntityErr, err)
	} else if exist {
		return e.Wrap(e.ResourceCreated, e.ResourceAlreadyExists)
	}

	cate := category.NewCategory(newName, cmd.SeoDesc)

	return c.cateRepo.Save(ctx, cate)
}
