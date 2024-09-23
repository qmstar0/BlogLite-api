package categories

import "context"

type CategoryRepository interface {
	Save(ctx context.Context, category *Category) error
	Find(ctx context.Context, slug string) (*Category, error)
	CheckNameExist(ctx context.Context, name string) (bool, error)
	Remove(ctx context.Context, category *Category) error
}
