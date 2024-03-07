package query

import (
	"blog/pkg/rediscache"
	"categorys/domain/category"
	"common/handler"
	"context"
)

type GetAllCategory struct {
	Page uint32
	Num  uint32
}

type GetAllCategoryResult struct {
	Cid         int
	Name        string
	DisplayName string
	SeoDesc     string
	Num         uint
}

type GetAllCategoryHandler handler.QueryHandler[GetAllCategory, GetAllCategoryResult]

type getAllCateogryHandler struct {
	cache    rediscache.Cache
	cateRepo category.CategoryRepository
}

func NewGetAllCategoryHandler(repo category.CategoryRepository, cache rediscache.Cache) GetAllCategoryHandler {
	return &getAllCateogryHandler{cateRepo: repo, cache: cache}
}

func (g getAllCateogryHandler) Handle(ctx context.Context, q GetAllCategory) (GetAllCategoryResult, error) {
	//TODO implement me
	panic("implement me")
}
