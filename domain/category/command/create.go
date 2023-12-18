package command

import "blog/domain/category/repository"

type CreateCategoryCommand struct {
	Name        string
	DisplayName string
	SeoDesc     string
}

type CreateCategoryCommandHandle struct {
	Repository repository.CategoryRepository
}

func (c CreateCategoryCommandHandle) CommandHandler() {

}
