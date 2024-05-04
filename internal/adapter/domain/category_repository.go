package domain

import (
	"context"
	"fmt"
	"github.com/uptrace/bun"
	"go-blog-ddd/internal/adapter/transaction"
	"go-blog-ddd/internal/domain/aggregates"
	"go-blog-ddd/internal/domain/values"
	"time"
)

type CategoryRepository struct {
	db            transaction.TransactionContext
	sequenceTable string
}

func NewCategoryRepository(db transaction.TransactionContext) CategoryRepository {
	if db == nil {
		panic("missing db")
	}
	repo := CategoryRepository{db: db, sequenceTable: "category_sequence"}
	repo.InitTable()

	return repo
}

func (r CategoryRepository) InitTable() {
	const InitTableDuration = time.Second * 10
	var err error

	timeout, cancelFunc := context.WithTimeout(context.Background(), InitTableDuration)
	defer cancelFunc()

	_, err = r.db.NewCreateTable().
		Model((*CategoryM)(nil)).
		IfNotExists().
		Exec(timeout)
	if err != nil {
		panic(err)
	}
	_, err = r.db.ExecContext(
		timeout,
		"CREATE SEQUENCE IF NOT EXISTS ? START 1 MINVALUE 1 MAXVALUE 10000",
		bun.Ident(r.sequenceTable),
	)
	if err != nil {
		panic(err)
	}
}

func (r CategoryRepository) DropTable() {
	const DropTableDuration = time.Second * 10
	timeout, cancelFunc := context.WithTimeout(context.Background(), DropTableDuration)
	defer cancelFunc()
	_, err := r.db.ExecContext(
		timeout,
		"DROP SEQUENCE ?",
		bun.Ident(r.sequenceTable),
	)
	if err != nil {
		fmt.Println("on drop table", err)
	}
	_, err = r.db.NewDropTable().Model((*CategoryM)(nil)).Exec(timeout)
	if err != nil {
		fmt.Println("on drop table", err)
	}
}

func (r CategoryRepository) NextID(ctx context.Context) uint32 {
	var nextID uint32
	rowContext := r.db.QueryRowContext(ctx, "select nextval('?');", bun.Ident(r.sequenceTable))
	err := rowContext.Scan(&nextID)
	if err != nil {
		panic(err)
	}
	return nextID
}

func (r CategoryRepository) FindByID(ctx context.Context, id uint32) (*aggregates.Category, error) {
	var cate = make([]*CategoryM, 0, 1)
	err := r.db.NewSelect().Model(&cate).Where("id = ?", id).Limit(1).Scan(ctx)
	if err != nil {
		return nil, err
	}
	if len(cate) >= 1 {
		return CategoryModelToAggregate(cate[0]), nil
	}
	return nil, nil
}

func (r CategoryRepository) FindByName(ctx context.Context, name values.CategoryName) (*aggregates.Category, error) {
	var cate = make([]*CategoryM, 0, 1)
	err := r.db.NewSelect().Model(&cate).Where("name = ?", name.String()).Limit(1).Scan(ctx)
	if err != nil {
		return nil, err
	}
	if len(cate) >= 1 {
		return CategoryModelToAggregate(cate[0]), nil
	}
	return nil, nil
}

func (r CategoryRepository) FindByIDOrErr(ctx context.Context, id uint32) (*aggregates.Category, error) {
	var model = &CategoryM{ID: id}
	err := r.db.NewSelect().Model(model).WherePK().Scan(ctx)
	if err != nil {
		return nil, err
	}
	return CategoryModelToAggregate(model), nil
}

func (r CategoryRepository) Save(ctx context.Context, category *aggregates.Category) error {
	model := CategoryAggregateToDBModel(category)
	return r.db.Transaction(ctx, func(tctx transaction.TransactionContext) error {
		ins := tctx.NewInsert().Model(model).On("CONFLICT (id) DO UPDATE")
		_, err := ins.Exec(tctx)
		return err
	})
}

func (r CategoryRepository) Delete(ctx context.Context, cid uint32) error {
	return r.db.Transaction(ctx, func(tctx transaction.TransactionContext) error {
		_, err := tctx.NewDelete().Model(&CategoryM{ID: cid}).WherePK().Exec(tctx)
		return err
	})
}
