package domain

import (
	"context"
	"go-blog-ddd/internal/adapter/transaction"
)

type TagReadModel struct {
	db transaction.TransactionContext
}

func NewTagReadModel(db transaction.TransactionContext) TagReadModel {
	return TagReadModel{db: db}
}
func (t TagReadModel) All(ctx context.Context) ([]string, error) {
	var result = make([]string, 0)
	err := t.db.NewSelect().Model((*PostTagM)(nil)).ColumnExpr("DISTINCT tag").Scan(ctx, &result)
	if err != nil {
		return nil, err
	}
	return result, nil
}
