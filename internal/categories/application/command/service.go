package command

import "context"

type GetCategoryUsedService interface {
	IsUsed(ctx context.Context, categorySlug string) (bool, error)
}
