package db

import (
	"blog/infrastructure/persistence/model"
	"time"
)

func Migrattion() error {
	if err := db.
		Set("gorm:table_option", "charset=utf8mb4").
		Set("gorm:query_options", map[string]any{"timeout": 5 * time.Second}).
		AutoMigrate(
			&model.CategoryM{},
			&model.UserM{},
		); err != nil {
		return err
	}

	return nil
}
