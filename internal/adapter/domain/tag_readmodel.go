package domain

import (
	"context"
	"go-blog-ddd/internal/adapter/e"
	"go-blog-ddd/internal/adapter/transaction"
	"go-blog-ddd/internal/application/query"
)

type TagReadModel struct {
	db transaction.TransactionContext
}

func NewTagReadModel(db transaction.TransactionContext) TagReadModel {
	return TagReadModel{db: db}
}

func (t TagReadModel) All(ctx context.Context) (query.TagListView, error) {
	var result = make([]*PostTagM, 0)
	err := t.db.NewSelect().Model(&result).
		ColumnExpr("tag, COUNT(*) as num").
		Group("tag").
		Scan(ctx)
	if err != nil {
		return query.TagListView{}, e.RErrDatabase.WithError(err)
	}

	return PostTagsModelToView(result), nil
}
