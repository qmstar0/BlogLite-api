package repository

import "blog/domain/post/entity"

type PostCategoryRepository interface {
	Save(category entity.PostCategoryImpl) error
	FindByPid(pid string) (entity.PostCategory, error)
}
