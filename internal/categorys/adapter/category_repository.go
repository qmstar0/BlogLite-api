package adapter

import (
	"categorys/domain/category"
	"context"
	"gorm.io/gorm"
)

type categoryRepositoryImpl struct {
	db *gorm.DB
}

func (c categoryRepositoryImpl) FindById(ctx context.Context, id string) (*category.Category, error) {
	return nil, nil
}

func (c categoryRepositoryImpl) Save(ctx context.Context, cate *category.Category) error {

	return nil
}

func NewCategoryRepository(db *gorm.DB) category.CategoryRepository {
	return &categoryRepositoryImpl{db: db}
}
