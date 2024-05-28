package domain

import (
	"context"
	"fmt"
	"github.com/qmstar0/nightsky-api/internal/domain/aggregates"
	"github.com/qmstar0/nightsky-api/internal/domain/values"
	"github.com/qmstar0/nightsky-api/internal/pkg/e"
	"github.com/qmstar0/nightsky-api/pkg/transaction"
	"github.com/uptrace/bun"

	"time"
)

type PostRepository struct {
	db            transaction.TransactionContext
	sequenceTable string
}

func NewPostRepository(db transaction.TransactionContext) PostRepository {
	if db == nil {
		panic("missing db")
	}
	repo := PostRepository{db: db, sequenceTable: "post_sequence"}
	repo.InitTable()

	return repo
}

func (p PostRepository) InitTable() {
	const InitTableDuration = time.Second * 10
	var err error

	timeout, cancelFunc := context.WithTimeout(context.Background(), InitTableDuration)
	defer cancelFunc()

	_, err = p.db.NewCreateTable().Model((*PostM)(nil)).IfNotExists().Exec(timeout)
	if err != nil {
		panic(err)
	}
	_, err = p.db.NewCreateTable().Model((*PostTagM)(nil)).IfNotExists().Exec(timeout)
	if err != nil {
		panic(err)
	}
	_, err = p.db.ExecContext(
		timeout,
		"CREATE SEQUENCE IF NOT EXISTS ? START 1000 INCREMENT BY 10 MINVALUE 1000",
		bun.Ident(p.sequenceTable),
	)
	if err != nil {
		panic(err)
	}
}

func (p PostRepository) DropTable() {
	const DropTableDuration = time.Second * 10
	timeout, cancelFunc := context.WithTimeout(context.Background(), DropTableDuration)
	defer cancelFunc()
	_, err := p.db.ExecContext(
		timeout,
		"DROP SEQUENCE ?",
		bun.Ident(p.sequenceTable),
	)
	if err != nil {
		fmt.Println("on drop table", err)
	}
	_, err = p.db.NewDropTable().Model((*PostM)(nil)).Exec(timeout)
	if err != nil {
		fmt.Println("on drop table", err)
	}
	_, err = p.db.NewDropTable().Model((*PostTagM)(nil)).Exec(timeout)
	if err != nil {
		fmt.Println("on drop table", err)
	}
}
func (p PostRepository) NextID(ctx context.Context) (uint32, error) {
	var nextID uint32
	rowContext := p.db.QueryRowContext(ctx, "select nextval('?');", bun.Ident(p.sequenceTable))
	err := rowContext.Scan(&nextID)
	if err != nil {
		return 0, e.RErrDatabase.WithError(err)
	}
	return nextID, nil
}

func (p PostRepository) FindByUri(ctx context.Context, uri values.PostUri) (*aggregates.Post, error) {
	var posts = make([]*PostM, 0, 1)
	err := p.db.NewSelect().
		Model(&posts).
		Relation("PostTags").
		Relation("Category").
		Where("post.uri = ?", uri.String()).
		Limit(1).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	if len(posts) <= 0 {
		return nil, nil
	}
	return PostModelToAggregate(posts[0]), nil
}

func (p PostRepository) FindOrErrByUri(ctx context.Context, uri values.PostUri) (*aggregates.Post, error) {
	found, err := p.FindByUri(ctx, uri)
	if err != nil {
		return nil, err
	}
	if found == nil {
		return nil, e.RErrResourceNotExists
	}
	return found, nil
}

func (p PostRepository) FindOrErrByID(ctx context.Context, id uint32) (*aggregates.Post, error) {
	var posts = make([]*PostM, 0, 1)
	err := p.db.NewSelect().
		Model(&posts).
		Relation("PostTags").
		Relation("Category").
		Where("post.id = ?", id).
		Limit(1).
		Scan(ctx)
	if err != nil {
		return nil, e.RErrDatabase.WithError(err)
	}
	if len(posts) <= 0 {
		return nil, e.RErrResourceNotExists
	}
	return PostModelToAggregate(posts[0]), nil
}

func (p PostRepository) Save(ctx context.Context, post *aggregates.Post) error {
	model := PostAggregateToModel(post)
	return p.db.Transaction(ctx, func(tctx transaction.TransactionContext) error {
		_, err := tctx.NewInsert().Model(model).On("CONFLICT (id) DO UPDATE").Exec(tctx)
		if err != nil {
			return e.RErrDatabase.WithError(err)
		}

		if len(model.PostTags) > 0 {
			_, err = tctx.NewDelete().Model(&PostTagM{PostID: post.ID}).WherePK("post_id").Exec(ctx)
			if err != nil {
				return e.RErrDatabase.WithError(err)
			}
			_, err = tctx.NewInsert().Model(&model.PostTags).Exec(tctx)
			if err != nil {
				return e.RErrDatabase.WithError(err)
			}
		}
		return nil
	})
}

func (p PostRepository) Delete(ctx context.Context, pid uint32) error {
	return p.db.Transaction(ctx, func(tctx transaction.TransactionContext) error {
		_, err := tctx.NewDelete().Model(&PostTagM{PostID: pid}).WherePK("post_id").Exec(ctx)
		if err != nil {
			return e.RErrDatabase.WithError(err)
		}
		_, err = tctx.NewDelete().Model(&PostM{ID: pid}).WherePK().Exec(tctx)
		if err != nil {
			return e.RErrDatabase.WithError(err)
		}
		return nil
	})
}

func (p PostRepository) ResourceUniquenessCheck(ctx context.Context, uri values.PostUri) error {
	exists, err := p.db.NewSelect().Model(&PostM{Uri: uri.String()}).
		WherePK("uri").
		Limit(1).
		Exists(ctx)
	if err != nil {
		return e.RErrDatabase.WithError(err)
	}
	if exists {
		return e.RErrResourceExists
	}
	return nil
}
