package command

import (
	"categorys/domain/category"
	"common/e"
	"common/handler"
	"context"
)

type UpdateCategory struct {
	Name        string
	DisplayName string
	SeoDesc     string
}

type UpdateCategoryHandler handler.CommandHandler[UpdateCategory]

type updateCategoryHandler struct {
	cateRepo category.CategoryRepository
}

func NewUpdataCategoryHandler(repo category.CategoryRepository) UpdateCategoryHandler {
	return &updateCategoryHandler{cateRepo: repo}
}

func (u updateCategoryHandler) Handle(ctx context.Context, cmd UpdateCategory) error {
	newName, err := category.NewName(cmd.Name)
	if err != nil {
		return e.Wrap(e.NewValueObjectErr, err)
	}
	cate, err := u.cateRepo.Find(ctx, newName.ToUint32())
	if err != nil {
		return e.Wrap(e.ResourceDoesNotExist, err)
	}

	cate.ChangeCategory(cmd.DisplayName, cmd.SeoDesc)

	return u.cateRepo.Save(ctx, cate)
}
