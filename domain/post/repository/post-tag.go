package repository

import "blog/domain/post/entity"

type PostTagRepository interface {
	Save(postTag entity.PostTagImpl) error
	FindByPid(pid string) (entity.PostTag, error)
}
