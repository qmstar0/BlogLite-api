package query

import (
	"context"
	"github.com/qmstar0/BlogLite-api/internal/common/constant"
	"github.com/qmstar0/BlogLite-api/pkg/utils"
)

type ArticleList struct {
	Category *string
	Tags     []string
	Page     *int
	Limit    *int
	Extra    bool
}

type ArticleListReadmodel interface {
	ArticleList(ctx context.Context, offset, limit int, tags []string, categoryID *string, extra bool) ([]ArticleView, error)
}

type ArticleListHandler struct {
	rm ArticleListReadmodel
}

func NewArticleListHandler(rm ArticleListReadmodel) *ArticleListHandler {
	return &ArticleListHandler{rm: rm}
}

func (a *ArticleListHandler) Handle(ctx context.Context, query ArticleList) (ArticleListView, error) {
	var (
		page  = 1 // 默认第一页
		limit = constant.ArticleListDefaultLimit
	)

	if query.Page != nil && *query.Page > 1 {
		page = *query.Page
	}

	if query.Limit != nil && *query.Limit > 0 {
		limit = *query.Limit
	}

	list, err := a.rm.ArticleList(
		ctx,
		utils.Offset(page, limit),
		limit+1,
		query.Tags,
		query.Category,
		query.Extra,
	)

	if err != nil {
		return ArticleListView{}, err
	}
	listLen := len(list)
	next := listLen > limit
	if next {
		list = list[:listLen-1]
	}
	return ArticleListView{
		Count: len(list),
		Page:  page,
		Items: list,
		Prev:  page > 1,
		Next:  next,
	}, nil
}
