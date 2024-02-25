package category

import "context"

type CategoryRepository interface {
	FindById(ctx context.Context, id string) (*Category, error)
	Save(ctx context.Context, cate *Category) error
}
