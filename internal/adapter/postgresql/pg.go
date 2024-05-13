package postgresql

import (
	"crypto/tls"
	"database/sql"
	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"go-blog-ddd/config"
	"time"
)

var db *bun.DB

func GetDB() *bun.DB {
	return db
}

func Init() (closeFn func() error) {
	var err error
	sqlDB := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithNetwork("tcp"),
		pgdriver.WithAddr(config.Cfg.Postgre.Addr),
		pgdriver.WithTLSConfig(&tls.Config{InsecureSkipVerify: true}),
		pgdriver.WithUser(config.Cfg.Postgre.User),
		pgdriver.WithPassword(config.Cfg.Postgre.Password),
		pgdriver.WithDatabase(config.Cfg.Postgre.Database),

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
