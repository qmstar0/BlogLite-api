package adapter

import (
	"context"
	"encoding/json"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/qmstar0/BlogLite-api/internal/articles/domain/articles"
	"github.com/qmstar0/BlogLite-api/internal/common/domain"
	"github.com/qmstar0/BlogLite-api/internal/common/e"
	"github.com/qmstar0/BlogLite-api/pkg/logging"
	"gorm.io/gorm"
	"sync"
)

type ArticleTagRelation struct {
	ID         uint32 `gorm:"primaryKey"`
	ArticleURI string
	Tag        string
}

func (a ArticleTagRelation) TableName() string {
	return "article_tag"
}

type PostgresArticleTagReadmodel struct {
	db   *gorm.DB
	lock *sync.Mutex
}

func NewPostgresArticleTagReadmodel(db *gorm.DB) *PostgresArticleTagReadmodel {
	if err := db.AutoMigrate(&ArticleTagRelation{}); err != nil {
		logging.Logger().Fatal("数据表初始化失败", "err", err)
	}
	return &PostgresArticleTagReadmodel{db: db, lock: &sync.Mutex{}}
}

func (p PostgresArticleTagReadmodel) Topics() []string {
	return []string{
		domain.EventTopic(articles.ArticleTagsModifiedEvent{}),
		domain.EventTopic(articles.ArticleDeletedEvent{}),
	}
}

func (p PostgresArticleTagReadmodel) Handle(msg *message.Message) ([]*message.Message, error) {
	p.lock.Lock()
	defer p.lock.Unlock()
	switch message.SubscribeTopicFromCtx(msg.Context()) {
	case "articles.ArticleTagsModifiedEvent":
		return nil, p.handleArticleTagsModifiedEvent(msg)
	case "articles.ArticleDeletedEvent":
		return nil, p.handleArticleDeletedEvent(msg)
	default:
		return nil, nil
	}
}

func (p PostgresArticleTagReadmodel) handleArticleTagsModifiedEvent(msg *message.Message) error {
	var event articles.ArticleTagsModifiedEvent
	if err := json.Unmarshal(msg.Payload, &event); err != nil {
		return err
	}
	ctx := msg.Context()
	return p.db.Transaction(func(tx *gorm.DB) error {
		if err := p.removeTagsByArticleURI(ctx, tx, event.URI); err != nil {
			return err
		}
		newTagsLen := len(event.NewTags)
		if newTagsLen <= 0 {
			return nil
		}
		var relation = make([]ArticleTagRelation, newTagsLen)
		for i, tag := range event.NewTags {
			relation[i] = ArticleTagRelation{ArticleURI: event.URI, Tag: tag}
		}

		if err := tx.WithContext(ctx).Create(&relation).Error; err != nil {
			return err
		}
		return nil
	})
}

func (p PostgresArticleTagReadmodel) handleArticleDeletedEvent(msg *message.Message) error {
	var event articles.ArticleDeletedEvent
	if err := json.Unmarshal(msg.Payload, &event); err != nil {
		return err
	}
	err := p.removeTagsByArticleURI(msg.Context(), p.db, event.URI)
	if err != nil {
		return err
	}
	return nil
}

func (p PostgresArticleTagReadmodel) removeTagsByArticleURI(ctx context.Context, tx *gorm.DB, uri string) error {
	return tx.WithContext(ctx).Where("article_uri = ?", uri).Delete(&ArticleTagRelation{}).Error
}

func (p PostgresArticleTagReadmodel) TagList(ctx context.Context) ([]string, error) {
	var tags = make([]string, 0)
	err := p.db.WithContext(ctx).
		Model(&ArticleTagRelation{}).
		Joins("LEFT JOIN article_metadata ON article_metadata.uri = article_tag.article_uri").
		Where("article_metadata.visibility = true").
		Distinct("tag").
		Pluck("tag", &tags).Error

	if err != nil {
		return nil, e.InternalServiceError(err.Error())
	}
	return tags, nil
}
