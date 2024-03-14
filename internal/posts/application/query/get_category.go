package query

import (
	"blog/pkg/rediscache"
	"categorys/domain/category"
	"common/handler"
	"context"
)

type GetCategory struct {
	Name string
}

type CategoryView struct {
	Cid         uint32 `json:"cid"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	SeoDesc     string `json:"seoDesc"`
	Num         uint   `json:"num"`
}

type GetCategoryHandler handler.QueryHandler[GetCategory, CategoryView]

type getCateogryHandler struct {
	cache    rediscache.Cache
	cateRepo category.CategoryRepository
}

func NewGetCategoryHandler(repo category.CategoryRepository, cache rediscache.Cache) GetCategoryHandler {
	return &getCateogryHandler{cateRepo: repo, cache: cache}
}

func (g getCateogryHandler) Handle(ctx context.Context, q GetCategory) (CategoryView, error) {
	panic("impl")
}
