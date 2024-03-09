package command

import (
	"categorys/domain/category"
	"common/e"
	"common/handler"
	"context"
)

type DeleteCategory struct {
	Name string
}

type DeleteCategoryHandler handler.CommandHandler[DeleteCategory]

type deleteCategoryHandler struct {
	cateRepo category.CategoryRepository
}

func NewDeleteCategoryHandler(repo category.CategoryRepository) DeleteCategoryHandler {
	return &deleteCategoryHandler{cateRepo: repo}
}

func (u deleteCategoryHandler) Handle(ctx context.Context, cmd DeleteCategory) error {
	newName, err := category.NewName(cmd.Name)
	if err != nil {
		return e.Wrap(e.NewValueObjectErr, err)
	}
	cate, err := u.cateRepo.Find(ctx, newName.ToID())
	if err != nil {
		return e.Wrap(e.ResourceDoesNotExist, err)
	}

	cate.Delete()

	return u.cateRepo.Save(ctx, cate)
}
