package command

import (
	"blog/pkg/cqrs"
	"categorys/domain/category"
	"categorys/domain/event"
	"common/handler"
	"context"
)

type CreateCategory struct {
	Name        string
	DisplayName string
	SeoDesc     string
}

type CreateCategoryHandler handler.CommandHandler[CreateCategory]

type createCategoryHandler struct {
	cateRepo category.CategoryRepository
	pub      cqrs.Publisher
}

func NewCreateCategoryHandler(repo category.CategoryRepository, pub cqrs.Publisher) CreateCategoryHandler {
	return &createCategoryHandler{
		cateRepo: repo,
		pub:      pub,
	}
}

func (c createCategoryHandler) Handle(ctx context.Context, cmd CreateCategory) error {
	var err error

	newName, err := category.NewName(cmd.Name)
	if err != nil {
		return err
	}

	cate := category.CreateCategory(newName, cmd.DisplayName, cmd.SeoDesc)

	err = c.pub.Publish(ctx, event.CategoryCreated{
		Cid:         cate.Cid,
		Name:        cate.Name.String(),
		DisplayName: cate.DisplayName,
		SeoDesc:     cate.SeoDesc,
	})
	if err != nil {
		return err
	}
	return c.cateRepo.Save(ctx, cate)
}
