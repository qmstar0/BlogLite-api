package domain

import (
	"context"
	"github.com/qmstar0/domain/internal/apps/query"
	"github.com/qmstar0/domain/internal/pkg/e"
	"github.com/qmstar0/domain/pkg/transaction"
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
