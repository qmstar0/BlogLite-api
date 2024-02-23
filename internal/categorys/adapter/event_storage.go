package adapter

import (
	"common/events"
	"context"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type CategoryEventStore struct {
	db *gorm.DB
}

func NewCategoryEventStore(db *gorm.DB) *CategoryEventStore {
	c := &CategoryEventStore{db: db}
	c.Migrattion()
	return c
}

func (c CategoryEventStore) Store(ctx context.Context, event events.Event) error {
	var (
		err error
		tx  = c.db.WithContext(ctx).Model(CategoryDomainEventStoreModel{}).Begin()
	)

	model, err := eventToModel(event)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Create(model).Error
}

func (c CategoryEventStore) Migrattion() {
	err := c.db.
		Set("gorm:table_option", "charset=utf8mb4").
		Set("gorm:query_options", map[string]any{"timeout": 5 * time.Second}).
		AutoMigrate(&CategoryDomainEventStoreModel{})
	if err != nil {
		panic(fmt.Errorf("err on AutoMigrate: %w", err))
	}
}
