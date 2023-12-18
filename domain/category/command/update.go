package command

import "blog/domain/category/repository"

type UpdateCommand struct {
	Id          uint
	ParendId    uint
	DisplayName string
	SeoDesc     string
}

type UpdadteCommandHandler struct {
	repository repository.CategoryRepository
}

func (c UpdadteCommandHandler) CommandHandler() {

}
