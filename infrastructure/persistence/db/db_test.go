package db_test

import (
	"blog/infrastructure/persistence/db"
	"testing"
)

func TestGetDB(t *testing.T) {
	getDB := db.GetDB()
	t.Log(getDB)
}
