package adapter

import (
	"context"
	"fmt"
	"github.com/qmstar0/BlogLite-api/internal/categories/domain/categories"
	"github.com/qmstar0/BlogLite-api/internal/common/e"
)

type CategoryValidityCheckService struct {
	repo categories.CategoryRepository
}

func NewCategoryValidityCheckService(repo categories.CategoryRepository) *CategoryValidityCheckService {
	return &CategoryValidityCheckService{repo: repo}
}

func (c CategoryValidityCheckService) CategoryExist(ctx context.Context, categoryID string) error {
	find, err := c.repo.Find(ctx, categoryID)
	if err != nil {
		return err
	}
	if find == nil {
		return e.ResourceDoesNotExist.WithMessage(fmt.Sprintf("%s: %s", "该分类不存在", categoryID))
	}
	return nil
}
