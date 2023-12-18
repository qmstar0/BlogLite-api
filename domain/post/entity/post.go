package entity

import (
	"blog/domain/post/valueobject"
)

type Post interface {
	Create()
	Update()
	Delete()
}

type PostImpl struct {
	Pid      string
	Title    string
	Slug     valueobject.PostSlug
	Summary  string
	Original string
	Content  string
	State    PostState
	Tags     []PostTag
	Category PostCategory
}
