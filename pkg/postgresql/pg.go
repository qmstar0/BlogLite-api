package postgresql

import (
	"crypto/tls"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"time"
)

var db *bun.DB

func GetDB() *bun.DB {
	return db
}

func Init(addr, user, pwd, database string) (closeFn func() error) {
	var err error
	sqlDB := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(addr),
		pgdriver.WithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
		pgdriver.WithUser(user),
		pgdriver.WithPassword(pwd),
		pgdriver.WithDatabase(database),

		pgdriver.WithTimeout(5*time.Second),
		pgdriver.WithDialTimeout(5*time.Second),
		pgdriver.WithReadTimeout(5*time.Second),
		pgdriver.WithWriteTimeout(5*time.Second),
	))

	sqlDB.SetConnMaxLifetime(time.Minute * 5)
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(50)

	db = bun.NewDB(sqlDB, pgdialect.New())

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return db.Close
}
