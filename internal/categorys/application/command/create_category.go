package command

import (
	"categorys/domain/category"
	"common/decorator"
	"context"
)

type CreateCategory struct {
	Uid         int
	Name        string
	DisplayName string
	SeoDesc     string
}

type CreateCategoryHandler decorator.CommandHandler[CreateCategory]

type createCategoryHandler struct {
	repo category.CategoryRepository
}

func NewCreateCategoryHandler(repo category.CategoryRepository) CreateCategoryHandler {
	return &createCategoryHandler{repo: repo}
}

func (c createCategoryHandler) Handle(ctx context.Context, cmd CreateCategory) error {
	var err error
	//defer func() {
	//	if err != nil {
	//		//发布事件
	//	}
	//}()

	newName, err := category.NewName(cmd.Name)
	if err != nil {
		return err
	}

	cate := category.CreateCategory(newName, cmd.DisplayName, cmd.SeoDesc)

	err = c.repo.Save(cate)
	if err != nil {
		return err
	}
	return nil
}
