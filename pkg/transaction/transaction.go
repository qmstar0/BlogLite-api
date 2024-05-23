package transaction

import (
	"context"
	"github.com/uptrace/bun"
	"time"
)

const txCtxKey = "TransactionContextKey"

type TransactionContext interface {
	bun.IDB
	context.Context
	Transaction(ctx context.Context, f func(tctx TransactionContext) error) error
}

type transactionContext struct {
	bun.IDB
	ctx context.Context
}

func (t transactionContext) Deadline() (deadline time.Time, ok bool) {
	return t.ctx.Deadline()
}

func (t transactionContext) Done() <-chan struct{} {
	return t.ctx.Done()
}

func (t transactionContext) Err() error {
	return t.ctx.Err()
}

func (t transactionContext) Value(key any) any {
	return t.ctx.Value(key)
}

func NewTransactionContext(db *bun.DB) TransactionContext {
	ctx := context.WithValue(context.Background(), txCtxKey, db)
	t := transactionContext{db, ctx}
	return t
}

func (t transactionContext) Transaction(ctx context.Context, f func(TransactionContext) error) error {
	idb, ok := ctx.Value(txCtxKey).(bun.IDB)
	if !ok {
		idb = t.IDB
	}
	tx, err := idb.BeginTx(ctx, nil)
	if err != nil {
		return err
	}

	defer func() {
		if p := recover(); p != nil {
			_ = tx.Rollback()
			panic(p)
		}
	}()

	t.IDB = tx
	if err = f(t); err != nil {
		_ = tx.Rollback()
		return err
	}

	return tx.Commit()
}
