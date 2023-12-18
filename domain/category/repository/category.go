package repository

import (
	"blog/domain/category/entity"
)

type CategoryRepository interface {
	New(cate entity.CategoryImpl) error
	Save(cate entity.CategoryImpl) error
	Delete(id uint) error
}
