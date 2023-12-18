package repository

import "blog/domain/post/entity"

type PostRepository interface {
	New(post entity.PostImpl) error
	Save(post entity.PostImpl) error
	FindByPid(pid string) (entity.Post, error)
	FindBySlug(Slug string) (entity.Post, error)
}
