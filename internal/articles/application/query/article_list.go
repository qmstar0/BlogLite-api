package query

import (
	"context"
	"github.com/qmstar0/BlogLite-api/internal/common/constant"
	"github.com/qmstar0/BlogLite-api/internal/common/e"
	"github.com/qmstar0/BlogLite-api/pkg/utils"
	"strings"
)

type ArticleList struct {
	Filter           *string
	Page             *int
	Limit            *int
	IncludeInvisible bool
}

type ArticleListReadmodel interface {
	ArticleList(ctx context.Context, offset, limit int, tags []string, categoryID *string, includeInvisible bool) ([]ArticleView, error)
}

type ArticleListHandler struct {
	rm ArticleListReadmodel
}

func NewArticleListHandler(rm ArticleListReadmodel) *ArticleListHandler {
	return &ArticleListHandler{rm: rm}
}

func (a *ArticleListHandler) Handle(ctx context.Context, query ArticleList) (ArticleListView, error) {
	var (
		page       = 1 // 默认第一页
		limit      = constant.ArticleListDefaultLimit
		tags       []string
		categoryID *string
	)

	if query.Filter != nil {
		split := strings.Split(strings.TrimSpace(*query.Filter), ";")
		if len(split) > 10 {
			return ArticleListView{}, e.InvalidParametersError("无效参数")
		}

		for _, s := range split {
			n := strings.SplitN(s, ":", 2)
			if len(n) <= 1 {
				continue
			}

			switch n[0] {
			case "category":
				categoryID = &n[1]
			case "tags":
				tags = strings.Split(n[1], ",")
			}
		}
	}

	if query.Page != nil && *query.Page > 1 {
		page = *query.Page
	}

	if query.Limit != nil && *query.Limit < 1 {
		limit = *query.Limit
	}

	list, err := a.rm.ArticleList(ctx, utils.Offset(page, limit), limit+1, tags, categoryID, query.IncludeInvisible)
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
