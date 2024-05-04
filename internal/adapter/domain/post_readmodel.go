package domain

import (
	"context"
	"github.com/uptrace/bun"
	"go-blog-ddd/internal/adapter/transaction"
	"go-blog-ddd/internal/application/query"
)

type PostReadModel struct {
	db transaction.TransactionContext
}

func NewPostReadModel(db transaction.TransactionContext) PostReadModel {
	return PostReadModel{db: db}
}
func (p PostReadModel) FindByID(ctx context.Context, pid uint32) (query.PostView, error) {
	var postM = &PostM{ID: pid}
	err := p.db.NewSelect().
		Model(postM).
		Relation("PostTags").
		Relation("Category").
		WherePK().
		Scan(ctx)
	if err != nil {
		return query.PostView{}, err
	}
	return PostModelToView(postM), nil
}

func (p PostReadModel) AllWithFilter(ctx context.Context, limit, offset int, tags []string, categroyID uint32) ([]query.PostView, error) {
	var postMS = make([]*PostM, 0)

	//tag filter
	var TagFilterFn func(*bun.SelectQuery) *bun.SelectQuery
	if len(tags) != 0 {
		filteredPost := p.db.NewSelect().
			Model((*PostTagM)(nil)).
			Column("post_id").
			Where("tag IN (?)", bun.In(tags)).
			Group("post_id").
			Having("COUNT(DISTINCT lang) >= ?")
		TagFilterFn = func(selectQuery *bun.SelectQuery) *bun.SelectQuery {
			return selectQuery.Where("post_id IN (?)", filteredPost)
		}
	}

	//category filter
	var CategoryFilterFn func(*bun.SelectQuery) *bun.SelectQuery
	if categroyID != 0 {
		CategoryFilterFn = func(selectQuery *bun.SelectQuery) *bun.SelectQuery {
			return selectQuery.Where("id = ?", categroyID)
		}
	}

	//main query
	selectQuery := p.db.NewSelect().
		Model(&postMS).
		Relation("PostTags", TagFilterFn).
		Relation("Category", CategoryFilterFn).
		Order("updated_at")
	if limit > 0 {
		selectQuery = selectQuery.Limit(limit)
	}
	if offset > 0 {
		selectQuery = selectQuery.Offset(offset)
	}

	if err := selectQuery.Scan(ctx); err != nil {
		return nil, err
	}

	postMLen := len(postMS)
	if postMLen <= 0 {
		return nil, nil
	}

	var result = make([]query.PostView, 0)
	for i, m := range postMS {
		result[i] = PostModelToView(m)
	}
	return result, nil
}
