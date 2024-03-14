package command

import (
	"categorys/domain/v1/category"
	"common/e"
	"common/handler"
	"context"
)

type DeleteCategory struct {
	ID uint32
}

type DeleteCategoryHandler handler.CommandHandler[DeleteCategory]

type deleteCategoryHandler struct {
	cateRepo category.CategoryRepository
}

func NewDeleteCategoryHandler(repo category.CategoryRepository) DeleteCategoryHandler {
	return &deleteCategoryHandler{cateRepo: repo}
}

func (u deleteCategoryHandler) Handle(ctx context.Context, cmd DeleteCategory) error {
	cate, err := u.cateRepo.Find(ctx, cmd.ID)
	if err != nil {
		return e.Wrap(e.FindEventErr, err)
	}

	cate.Delete()

	return u.cateRepo.Save(ctx, cate)
}
