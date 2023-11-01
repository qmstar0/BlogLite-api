package mysql

import (
	"time"

	"blog/infra/repository/model"
)

func Migrattion() error {
	if err := db.
		Set("gorm:table_option", "charset=utf8mb4").
		Set("gorm:query_options", map[string]any{"timeout": 5 * time.Second}).
		AutoMigrate(
			&model.User{},
			&model.Comments{},
			&model.ArticleMate{},
			&model.ArticleTags{},
			&model.ArticleCate{},
		); err != nil {
		return err
	}

	return nil
}
