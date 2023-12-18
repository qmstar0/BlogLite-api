package repository

import "blog/domain/post/entity"

type PostStateRepository interface {
	Save(post entity.PostStateImpl) error
	FindByPid(pid string) (entity.PostState, error)
}
