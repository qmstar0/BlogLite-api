package postgresql_test

import (
	"context"
	"fmt"
	"github.com/uptrace/bun"
	"go-blog-ddd/internal/adapter/postgresql"
	"go-blog-ddd/internal/adapter/transaction"
	"testing"
)

type Profile struct {
	ID     int64 `bun:",pk,autoincrement"`
	Lang   string
	Active bool
	Num    int
	UserID int64
}

// User has many profiles.
type User struct {
	ID       int64 `bun:",pk,autoincrement"`
	Name     string
	Age      int
	Tag      []string   `bun:",array"`
	Profiles []*Profile `bun:"rel:has-many,join:id=user_id"`
}

func TestName(t *testing.T) {

	fn := postgresql.Init()
	defer fn()
	db := postgresql.GetDB()
	tctx := transaction.NewTransactionContext(db)
	var err error
	ctx := context.Background()

	createSchema(ctx, tctx)
	users := []User{
		{
			//Name: "user 1",
			//Tag: []string{"log"},
		},
	}
	err = tctx.Transaction(ctx, func(tctx transaction.TransactionContext) error {
		//user := make([]*User, 0)
		query := tctx.NewSelect().
			Model(&users).
			//Column("user.*").
			Relation("Profiles", func(query *bun.SelectQuery) *bun.SelectQuery {
				tags := []string{"en", "md"}

				filteredUser := tctx.NewSelect().Model((*Profile)(nil)).Column("user_id").
					Where("lang IN (?)", bun.In(tags)).
					Group("user_id").
					Having("COUNT(DISTINCT lang) = ?", len(tags))
				return query.Where("user_id IN (?)", filteredUser)
			})
		//WherePK("tag").

		//Where("tag @> ")

		t.Log(query.String())
		//OrderExpr("user.id ASC").
		//Limit(1).
		return query.Scan(ctx)
	})
	if err != nil {
		t.Fatal(err)
	}
	for i, u := range users {
		fmt.Println(i, u.ID, u.Name, u.Age, u.Tag)
		for _, profile := range u.Profiles {
			fmt.Println("    ", profile)
		}
	}

}

func createSchema(ctx context.Context, db bun.IDB) {
	models := []interface{}{
		(*User)(nil),
		(*Profile)(nil),
	}
	for _, model := range models {
		if _, err := db.NewDropTable().Model(model).Exec(ctx); err != nil {
			panic(err)
		}
	}
	for _, model := range models {
		if _, err := db.NewCreateTable().Model(model).IfNotExists().Exec(ctx); err != nil {
			panic(err)
		}
	}
	//
	users := []*User{
		{ID: 1, Name: "user 1", Age: 12, Tag: make([]string, 0)},
		{ID: 2, Name: "user 2", Age: 12, Tag: []string{"test"}},
		{ID: 3, Name: "user 1", Age: 15, Tag: []string{"log"}},
		{ID: 4, Name: "admin", Age: 16, Tag: []string{"log", "test"}},
		{ID: 5, Name: "admin", Age: 20, Tag: []string{"log", "python"}},
		{ID: 6, Name: "root", Age: 30, Tag: make([]string, 0)},
	}
	if _, err := db.NewInsert().Model(&users).Exec(ctx); err != nil {
		panic(err)
	}

	profiles := []*Profile{
		{ID: 1, Lang: "en", Active: true, UserID: 1},
		{ID: 2, Lang: "ru", Active: true, UserID: 1},
		{ID: 3, Lang: "md", Active: false, UserID: 1},
		{ID: 7, Lang: "md", Active: false, UserID: 1},
		{ID: 4, Lang: "en", Active: false, UserID: 2},
		{ID: 5, Lang: "md", Active: false, UserID: 2},
		{ID: 6, Lang: "xx", Active: false, UserID: 2},
	}
	if _, err := db.NewInsert().Model(&profiles).Exec(ctx); err != nil {
		panic(err)
	}

}
