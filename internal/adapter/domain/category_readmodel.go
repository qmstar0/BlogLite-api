package domain

import (
	"context"
	"go-blog-ddd/internal/adapter/e"
	"go-blog-ddd/internal/adapter/transaction"
	"go-blog-ddd/internal/application/query"
)

type CategoryReadModel struct {
	db transaction.TransactionContext
}

func NewCategoryReadModel(db transaction.TransactionContext) CategoryReadModel {
	return CategoryReadModel{db: db}
}

func (q CategoryReadModel) All(ctx context.Context) (query.CategoryListView, error) {
	CategoryCount := q.db.NewSelect().
		Model((*PostM)(nil)).
		ColumnExpr("category_id, COUNT(*) AS category_count").
		Group("category_id")

	var categorys = make([]*CategoryM, 0)
	err := q.db.NewSelect().
		With("post_category", CategoryCount).
		Model(&categorys).
		ColumnExpr("category.*, post_category.category_count as num").
		Join("left join post_category on category.id = post_category.category_id").
		Order("id").
		Scan(ctx)
	if err != nil {
		return query.CategoryListView{}, e.RErrDatabase.WithError(err)
	}

	return CategoryModelToListView(categorys), nil
}
