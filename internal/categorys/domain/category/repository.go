package category

import "context"

type CategoryRepository interface {
	Save(ctx context.Context, cate *Category) error
}
