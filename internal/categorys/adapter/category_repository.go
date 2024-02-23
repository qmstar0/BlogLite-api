package adapter

import (
	"categorys/domain/category"
	"context"
)

type categoryRepositoryImpl struct {
	Map map[int]*category.Category
}

func (c categoryRepositoryImpl) Save(ctx context.Context, cate *category.Category) error {
	c.Map[cate.Cid] = cate
	return nil
}

func NewCategoryRepository() category.CategoryRepository {
	return &categoryRepositoryImpl{
		Map: make(map[int]*category.Category),
	}
}
