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

//func (p PostRepository) FindByTags(ctx context.Context, tagsVO []values.Tag) ([]*aggregates.Post, error) {
//	var postMs = make([]*PostM, 1)
//	err := p.db.NewSelect().Model((*PostM)(nil)).Where("tags @> ?", tagsVO).Scan(ctx, &postMs)
//	if err != nil {
//		return nil, err
//	}
//
//	var posts = make([]*aggregates.Post, len(postMs))
//	for i, postM := range postMs {
//		posts[i] = PostModelToAggregate(postM)
//	}
//	return posts, nil
//}

func (p PostRepository) FindByUri(ctx context.Context, uri values.PostUri) (*aggregates.Post, error) {
	var postM = make([]*PostM, 1)
	err := p.db.NewSelect().Model(&postM).Where("uri = ?", uri.String()).Limit(1).Scan(ctx)
	if err != nil {
		return nil, err
	}
	if len(postM) >= 1 {
		return PostModelToAggregate(postM[0]), nil
	}
	return nil, nil
}

func (p PostRepository) FindOrErrByUri(ctx context.Context, uri values.PostUri) (*aggregates.Post, error) {
	var postM = &PostM{Uri: uri.String()}
	err := p.db.NewSelect().Model(postM).WherePK("uri").Scan(ctx)
	if err != nil {
		return nil, err
	}
	return PostModelToAggregate(postM), nil
}

func (p PostRepository) FindOrErrByID(ctx context.Context, id uint32) (*aggregates.Post, error) {
	var postM = &PostM{ID: id}
	err := p.db.NewSelect().Model(postM).WherePK().Scan(ctx)
	if err != nil {
		return nil, err
	}
	return PostModelToAggregate(postM), nil
}

func (p PostRepository) Save(ctx context.Context, post *aggregates.Post) error {
	model := PostAggregateToModel(post)
	return p.db.Transaction(ctx, func(tctx transaction.TransactionContext) error {
		_, err := tctx.NewInsert().Model(model).On("CONFLICT (id) DO UPDATE").Exec(tctx)
		if err != nil {
			return err
		}
		if len(model.PostTags) > 0 {
			exec := tctx.NewInsert().Model(&model.PostTags).On("CONFLICT (id) DO UPDATE")
			_, err = exec.Exec(tctx)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func (p PostRepository) NextID(ctx context.Context) uint32 {
	var nextID uint32
	rowContext := p.db.QueryRowContext(ctx, "select nextval('?');", bun.Ident(p.sequenceTable))
	err := rowContext.Scan(&nextID)
	if err != nil {
		panic(err)
	}
	return nextID
}

func (p PostRepository) Delete(ctx context.Context, pid uint32) error {
	return p.db.Transaction(ctx, func(tctx transaction.TransactionContext) error {
		_, err := tctx.NewDelete().Model(&PostM{ID: pid}).WherePK().Exec(tctx)
		return err
	})
}
