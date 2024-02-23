package postgresql_test

import (
	_ "blog/pkg/env"
	"blog/pkg/postgresql"
	"testing"
)

func TestGetDB(t *testing.T) {
	db := postgresql.GetDB()
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatal(err)
	}
	err = sqlDB.Ping()
	if err != nil {
		t.Fatal(err)
	}
}
