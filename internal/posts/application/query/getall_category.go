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

type AllCategoryView []CategoryView

type GetAllCategoryHandler handler.QueryHandler[GetAllCategory, AllCategoryView]

type getAllCateogryHandler struct {
	cache    rediscache.Cache
	cateRepo category.CategoryRepository
}

func NewGetAllCategoryHandler(repo category.CategoryRepository, cache rediscache.Cache) GetAllCategoryHandler {
	return &getAllCateogryHandler{cateRepo: repo, cache: cache}
}

func (g getAllCateogryHandler) Handle(ctx context.Context, q GetAllCategory) (AllCategoryView, error) {
	//TODO implement me
	panic("implement me")
}
