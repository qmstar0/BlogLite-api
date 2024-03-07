package category

import "context"

type CategoryRepository interface {
	Find(ctx context.Context, aggID uint32) (*Category, error)
	Save(ctx context.Context, cate *Category) error
}
