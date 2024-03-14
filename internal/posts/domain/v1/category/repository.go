package category

import "context"

type CategoryRepository interface {
	Exist(ctx context.Context, id uint32) (bool, error)
	Find(ctx context.Context, id uint32) (*Category, error)
	Save(ctx context.Context, category *Category) error
}
