package adapter

import (
	"context"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/lib/pq"
	"github.com/qmstar0/BlogLite-api/internal/articles/domain/articles"
	"github.com/qmstar0/BlogLite-api/internal/common/domain"
	"github.com/qmstar0/BlogLite-api/internal/common/e"
	"github.com/qmstar0/BlogLite-api/pkg/logging"
	"gorm.io/gorm"
)

type Article struct {
	URI            string         `gorm:"primaryKey"`
	VersionList    pq.StringArray `gorm:"type:text[]"`
	CategoryID     string
	Tags           pq.StringArray `gorm:"type:text[]"`
	Visitility     bool
	CurrentVersion string
}

func (a Article) TableName() string {
	return "article"
}

type PostgresArticleRepository struct {
	db        *gorm.DB
	publisher message.Publisher
}

func NewPostgresArticleRepository(db *gorm.DB, publisher message.Publisher) *PostgresArticleRepository {
	if err := db.AutoMigrate(&Article{}); err != nil {
		logging.Logger().Fatal("数据表初始化失败", "err", err)
	}
	return &PostgresArticleRepository{db: db, publisher: publisher}
}

func (p PostgresArticleRepository) Find(ctx context.Context, uri articles.URI) (*articles.Article, error) {
	var article Article
	result := p.db.WithContext(ctx).Where("uri = ?", uri.String()).Limit(1).Find(&article)
	if result.Error != nil {
		return nil, e.InternalServiceError(result.Error.Error())
	} else if result.RowsAffected != 1 {
		return nil, nil
	}

	return articles.UnmarshalArticleFromDatabase(
		article.URI,
		article.Visitility,
		article.Tags,
		article.CategoryID,
		article.CurrentVersion,
		article.VersionList,
	)
}

func (p PostgresArticleRepository) Save(ctx context.Context, article *articles.Article) error {
	if err := p.db.WithContext(ctx).Where("uri = ?", article.Uri().String()).Save(&Article{
		URI:            article.Uri().String(),
		VersionList:    article.VersionList(),
		CategoryID:     article.CategoryID(),
		Tags:           article.TagGroup().Value(),
		Visitility:     article.IsVisible(),
		CurrentVersion: article.CurrentVersion(),
	}).Error; err != nil {
		return e.InternalServiceError(err.Error())
	}

	return p.Publish(article.Events())
}

func (p PostgresArticleRepository) Remove(ctx context.Context, article *articles.Article) error {
	if err := p.db.WithContext(ctx).Where("uri = ?", article.Uri().String()).Delete(&Article{}).Error; err != nil {
		return e.InternalServiceError(err.Error())
	}

	return p.Publish(article.Events())
}

func (p PostgresArticleRepository) UpdateArticle(ctx context.Context, uri articles.URI, updateFn func(*articles.Article) (*articles.Article, error)) error {
	found, err := p.Find(ctx, uri)
	if err != nil {
		return err
	}
	if found == nil {
		return e.ResourceDoesNotExist
	}
	article, err := updateFn(found)
	if err != nil {
		return err
	}
	return p.Save(ctx, article)
}

func (p PostgresArticleRepository) Publish(events []*domain.DomainEvent) error {
	for _, event := range events {
		err := p.publisher.Publish(event.Topic, message.NewMessage(event.EventID, event.Payload))
		if err != nil {
			return e.InternalServiceError(err.Error())
		}
	}
	return nil
}
