package adapter

import (
	"context"
	"github.com/qmstar0/BlogLite-api/internal/common/e"
	"gorm.io/gorm"
)

type GetCategoryUsedService struct {
	db *gorm.DB
}

func NewGetCategoryUsedService(db *gorm.DB) *GetCategoryUsedService {
	return &GetCategoryUsedService{db: db}
}

func (g GetCategoryUsedService) IsUsed(ctx context.Context, categorySlug string) (bool, error) {
	var count int64
	err := g.db.WithContext(ctx).
		Table("article_detail").
		Where("category_id = ?", categorySlug).
		Limit(1).
		Count(&count).Error
	if err != nil {
		return false, e.InternalServiceError(err.Error())
	}
	return count == 0, nil
}
