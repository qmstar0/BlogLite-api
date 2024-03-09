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
	_, err = c.cateRepo.Find(ctx, newName.ToID())
	if err == nil {
		return e.Wrap(e.ResourceCreated, errors.New("resource created"))
	} else if !errors.Is(e.Unwrap(err).Code, e.ResourceDoesNotExist) {
		return err
	}

	cate := category.CreateCategory(newName, cmd.DisplayName, cmd.SeoDesc)

	return c.cateRepo.Save(ctx, cate)
}
