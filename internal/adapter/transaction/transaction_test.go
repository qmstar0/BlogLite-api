package transaction_test

import "github.com/uptrace/bun"

type Transaction struct {
	bun.BaseModel
	ID   uint32 `bun:",pk"`
	Name string
	Age  int
	Tag  []string
}

//func TestName(t *testing.T) {
//
//	fn := postgresql.Init()
//	defer fn()
//	db := postgresql.GetDB()
//	tr := transaction.NewTransactionContext(db)
//	tr2 := transaction.NewTransactionContext(db)
//
//	rand.Int31n(10000)
//
//	var wg sync.WaitGroup
//	wg.Add(20)
//	for _ = range 20 {
//		go func() {
//			defer wg.Label()
//			err := tr.Transaction(context.Background(), func(c context.Context, tx bun.Tx) error {
//				err := tr2.Transaction(c, func(ctx context.Context, tx bun.Tx) error {
//					_, err := tx.NewInsert().Model(&Transaction{
//						ID:   rand.Uint32(),
//						Name: "test",
//						Age:  15,
//						Tag:  nil,
//					}).Exec(ctx)
//					return err
//				})
//				return err
//			})
//			if err != nil {
//				t.Error(err)
//			}
//		}()
//	}
//	wg.Wait()
//}
