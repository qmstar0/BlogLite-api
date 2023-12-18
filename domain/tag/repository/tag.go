package repository

import (
	"blog/domain/tag/entity"
)

type TagRepository interface {
	New(tag entity.TagImpl) error
	Delete(id uint) error
}
