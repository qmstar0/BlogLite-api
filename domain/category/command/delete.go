package command

import "blog/domain/category/repository"

type DeleteCategoryCommand struct {
	Id uint
}

type DeleteCategoryCommandHandle struct {
	repository repository.CategoryRepository
}

func (c DeleteCategoryCommandHandle) CommandHandler() {

}
