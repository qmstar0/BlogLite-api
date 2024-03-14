package command

import (
	"categorys/domain/v1/category"
	"common/e"
	"common/handler"
	"context"
)

type UpdateCategory struct {
	CategoryID uint32
	SeoDesc    string
}

type UpdateCategoryHandler handler.CommandHandler[UpdateCategory]

type updateCategoryHandler struct {
	cateRepo category.CategoryRepository
}

func NewUpdataCategoryHandler(repo category.CategoryRepository) UpdateCategoryHandler {
	return &updateCategoryHandler{cateRepo: repo}
}

func (u updateCategoryHandler) Handle(ctx context.Context, cmd UpdateCategory) error {
	cate, err := u.cateRepo.Find(ctx, cmd.CategoryID)
	if err != nil {
		return e.Wrap(e.FindEventErr, err)
	}

	cate.ModifySeoDesc(cmd.SeoDesc)

	return u.cateRepo.Save(ctx, cate)
}
