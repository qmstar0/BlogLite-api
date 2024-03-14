package post

import "context"

type PostRepository interface {
	Exist(ctx context.Context, id uint32) (bool, error)
	Find(ctx context.Context, id uint32) (*Post, error)
	Save(ctx context.Context, category *Post) error
}
