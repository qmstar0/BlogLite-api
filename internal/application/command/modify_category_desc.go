package command

import (
	"context"
	"go-blog-ddd/internal/domain/aggregates"
	"go-blog-ddd/internal/domain/commands"
)

type ModifyCategoryDescHandler struct {
	cateRepo aggregates.CategoryRepository
}

func NewModifyCategoryDescHandler(repo aggregates.CategoryRepository) ModifyCategoryDescHandler {
	if repo == nil {
		panic("missing CategoryRepository")
	}
	return ModifyCategoryDescHandler{cateRepo: repo}
}

func (u ModifyCategoryDescHandler) Handle(ctx context.Context, cmd commands.ModifyCategoryDesc) error {
	cate, err := u.cateRepo.FindOrErrByID(ctx, cmd.ID)
	if err != nil {
		return err
	}
	cate.SeoSeoDesc(cmd.NewDesc)

	return u.cateRepo.Save(ctx, cate)
}
