package postgresql

import (
	"database/sql"
	"github.com/charmbracelet/log"
	_ "github.com/lib/pq"
	"github.com/qmstar0/shutdown"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"time"
)

var db *bun.DB

func GetDB() *bun.DB {
	return db
}

func Init(dsn string) (closeFn func() error) {
	var err error

	sqlDB, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Error(err)
		shutdown.Exit(1)
	}
	sqlDB.SetConnMaxLifetime(time.Minute * 5)
	sqlDB.SetMaxIdleConns(50)
	sqlDB.SetMaxOpenConns(50)

	db = bun.NewDB(sqlDB, pgdialect.New())

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return db.Close
}
