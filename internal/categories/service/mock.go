package service

import "context"

type MockGetCategoryUsedService struct {
}

func (m MockGetCategoryUsedService) IsUsed(ctx context.Context, categorySlug string) (bool, error) {
	return false, nil
}
