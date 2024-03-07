package query

import (
	"blog/pkg/rediscache"
	"categorys/domain/category"
	"common/e"
	"common/handler"
	"common/idtools"
	"context"
)

type GetCategory struct {
	Name string
}

type GetCategoryResult struct {
	Cid         uint32 `json:"cid"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	SeoDesc     string `json:"seoDesc"`
	Num         uint   `json:"num"`
}

type GetCategoryHandler handler.QueryHandler[GetCategory, GetCategoryResult]

type getCateogryHandler struct {
	cache    rediscache.Cache
	cateRepo category.CategoryRepository
}

func NewGetCategoryHandler(repo category.CategoryRepository, cache rediscache.Cache) GetCategoryHandler {
	return &getCateogryHandler{cateRepo: repo, cache: cache}
}

func (g getCateogryHandler) Handle(ctx context.Context, q GetCategory) (GetCategoryResult, error) {
	cate, err := g.cateRepo.Find(ctx, idtools.NewHashID(q.Name))
	if err != nil {
		return GetCategoryResult{}, e.Wrap(e.ResourceDoesNotExist, err)
	}
	return GetCategoryResult{
		Cid:         cate.Cid,
		Name:        cate.Name.String(),
		DisplayName: cate.DisplayName,
		SeoDesc:     cate.SeoDesc,
		Num:         0,
	}, nil
}
