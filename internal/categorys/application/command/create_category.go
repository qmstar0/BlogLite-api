package command

import (
	"categorys/domain/category"
	"common/e"
	"common/handler"
	"context"
	"errors"
)

type CreateCategory struct {
	Name        string
	DisplayName string
	SeoDesc     string
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
	_, err = c.cateRepo.Find(ctx, newName.ToUint32())
	if err != nil {
		if e.Unwrap(err).Code != e.FindResultIsNull {
			return err
		}
	} else {
		return e.Wrap(e.ResourceCreated, errors.New("resource created"))
	}

	cate := category.CreateCategory(newName, cmd.DisplayName, cmd.SeoDesc)

	return c.cateRepo.Save(ctx, cate)
}
