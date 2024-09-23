package adapter

import (
	"context"
	"encoding/json"
	"github.com/ThreeDotsLabs/watermill/message"
	"github.com/lib/pq"
	"github.com/qmstar0/BlogLite-api/internal/articles/application/query"
	"github.com/qmstar0/BlogLite-api/internal/articles/domain/articles"
	"github.com/qmstar0/BlogLite-api/internal/categories/adapter"
	"github.com/qmstar0/BlogLite-api/internal/common/domain"
	"github.com/qmstar0/BlogLite-api/internal/common/e"
	"github.com/qmstar0/BlogLite-api/pkg/logging"
	"gorm.io/gorm"
	"sync"
	"time"
)

type ArticleDetail struct {
	URI                   string `gorm:"primaryKey"`
	CurrentVersion        string
	Visitility            bool
	FirstVersionCreatedAt time.Time
	CategoryID            string
}

func (a ArticleDetail) TableName() string {
	return "article_detail"
}

type articleDetailWithVersion struct {
	ArticleDetail
	ArticleVersion
	adapter.Category
	Tags pq.StringArray `gorm:"type:text[];not null"`
}

type PostgresArticleDetailReadmodel struct {
	db   *gorm.DB
	lock *sync.Mutex
}

func NewPostgresArticleDetailReadmodel(db *gorm.DB) *PostgresArticleDetailReadmodel {
	if err := db.AutoMigrate(&ArticleDetail{}); err != nil {
		logging.Logger().Fatal("数据表初始化失败", "err", err)
	}
	return &PostgresArticleDetailReadmodel{db: db, lock: &sync.Mutex{}}
}

func (p PostgresArticleDetailReadmodel) Topics() []string {
	return []string{
		domain.EventTopic(articles.ArticleDeletedEvent{}),
		domain.EventTopic(articles.ArticleContentSetSuccessfullyEvent{}),
		domain.EventTopic(articles.ArticleInitializedSuccessfullyEvent{}),
		domain.EventTopic(articles.ArticleVisibilityChangedEvent{}),
		domain.EventTopic(articles.ArticleFirstVersionCreatedEvent{}),
		domain.EventTopic(articles.ArticleCategoryChangedEvent{}),
	}
}

func (p PostgresArticleDetailReadmodel) Handle(msg *message.Message) ([]*message.Message, error) {
	p.lock.Lock()
	defer p.lock.Unlock()

	switch message.SubscribeTopicFromCtx(msg.Context()) {
	case "articles.ArticleDeletedEvent":
		return nil, p.handleArticleDeletedEvent(msg)
	case "articles.ArticleContentSetSuccessfullyEvent":
		return nil, p.handleArticleContentSetSuccessfullyEvent(msg)
	case "articles.ArticleInitializedSuccessfullyEvent":
		return nil, p.handleArticleInitializedSuccessfullyEvent(msg)
	case "articles.ArticleVisibilityChangedEvent":
		return nil, p.handleArticleVisibilityChangedEvent(msg)
	case "articles.ArticleFirstVersionCreatedEvent":
		return nil, p.handleArticleFirstVersionCreatedEvent(msg)
	case "articles.ArticleCategoryChangedEvent":
		return nil, p.handleArticleCategoryChangedEvent(msg)
	default:
		return nil, nil
	}
}

func (p PostgresArticleDetailReadmodel) handleArticleCategoryChangedEvent(msg *message.Message) error {
	var event articles.ArticleCategoryChangedEvent
	if err := json.Unmarshal(msg.Payload, &event); err != nil {
		return err
	}

	err := p.db.WithContext(msg.Context()).
		Model(&ArticleDetail{}).
		Where("uri = ?", event.URI).
		Update("category_id", event.NewCategoryID).Error
	if err != nil {
		return err
	}
	return nil
}

func (p PostgresArticleDetailReadmodel) handleArticleFirstVersionCreatedEvent(msg *message.Message) error {
	var event articles.ArticleFirstVersionCreatedEvent
	if err := json.Unmarshal(msg.Payload, &event); err != nil {
		return err
	}
	err := p.db.WithContext(msg.Context()).
		Model(&ArticleDetail{}).
		Where("uri = ?", event.URI).
		Update("first_version_created_at", event.CreatedAt).
		Update("current_version", event.Version).Error
	if err != nil {
		return err
	}
	return nil
}

func (p PostgresArticleDetailReadmodel) handleArticleDeletedEvent(msg *message.Message) error {
	var event articles.ArticleDeletedEvent
	if err := json.Unmarshal(msg.Payload, &event); err != nil {
		return err
	}

	if err := p.db.WithContext(msg.Context()).Where("uri = ?", event.URI).Delete(&ArticleDetail{}).Error; err != nil {
		return err
	}
	return nil
}

func (p PostgresArticleDetailReadmodel) handleArticleContentSetSuccessfullyEvent(msg *message.Message) error {
	var event articles.ArticleContentSetSuccessfullyEvent
	if err := json.Unmarshal(msg.Payload, &event); err != nil {
		return err
	}

	err := p.db.WithContext(msg.Context()).
		Model(&ArticleDetail{}).
		Where("uri = ?", event.URI).
		Update("current_version", event.Version).Error
	if err != nil {
		return err
	}
	return nil
}

func (p PostgresArticleDetailReadmodel) handleArticleInitializedSuccessfullyEvent(msg *message.Message) error {
	var event articles.ArticleInitializedSuccessfullyEvent
	if err := json.Unmarshal(msg.Payload, &event); err != nil {
		return err
	}

	err := p.db.WithContext(msg.Context()).Where("uri = ?", event.URI).Save(&ArticleDetail{
		URI:            event.URI,
		CurrentVersion: "",
		Visitility:     false,
		CategoryID:     event.CategoryID,
	}).Error
	if err != nil {
		return err
	}
	return nil
}

func (p PostgresArticleDetailReadmodel) handleArticleVisibilityChangedEvent(msg *message.Message) error {
	var event articles.ArticleVisibilityChangedEvent
	if err := json.Unmarshal(msg.Payload, &event); err != nil {
		return err
	}

	err := p.db.WithContext(msg.Context()).Model(&ArticleDetail{}).
		Where("uri = ?", event.URI).Update("visitility", event.Visibility).Error
	if err != nil {
		return err
	}
	return nil
}

func (p PostgresArticleDetailReadmodel) joinsQueryArticle(ctx context.Context) *gorm.DB {
	return p.db.WithContext(ctx).
		Model(&ArticleDetail{}).
		Select("article_detail.*, category.slug, category.name, article_version.*, array_remove(array_agg(article_tag.tag), null) AS tags").
		Joins("JOIN article_version ON article_version.uri = article_detail.uri").
		Joins("LEFT JOIN article_tag ON article_detail.uri = article_tag.article_uri").
		Joins("LEFT JOIN category ON article_detail.category_id = category.slug").
		Group("article_detail.uri, category.slug, article_version.version").
		Where("article_detail.visitility = true")
}

func (p PostgresArticleDetailReadmodel) ArticleDetail(ctx context.Context, uri string, version *string) (query.ArticleView, error) {
	tx := p.joinsQueryArticle(ctx).
		Limit(1).
		Where("article_detail.uri = ?", uri)

	if version != nil && *version != "" {
		tx = tx.Where("article_version.version = ?", *version)
	} else {
		tx = tx.Where("article_version.version = article_detail.current_version")
	}

	var article articleDetailWithVersion
	result := tx.Find(&article)
	if result.Error != nil {
		return query.ArticleView{}, e.InternalServiceError(result.Error.Error())
	} else if result.RowsAffected != 1 {
		return query.ArticleView{}, e.ResourceDoesNotExist
	}

	return query.ArticleView{
		Uri:         article.ArticleDetail.URI,
		Title:       article.Title,
		Description: article.ArticleVersion.Description,
		Note:        "",
		Content:     article.Content,
		CreatedAt:   article.CreatedAt.UnixMilli(),
		Category: query.ArticleCategory{
			Slug: article.Category.Slug,
			Name: article.Category.Name,
		},
		Tags: article.Tags,
	}, nil
}

func (p PostgresArticleDetailReadmodel) ArticleList(ctx context.Context, offset, limit int, tags []string, categoryID *string) ([]query.ArticleView, error) {
	tx := p.joinsQueryArticle(ctx).Offset(offset).Limit(limit).Where("article_version.version = article_detail.current_version")
	if len(tags) != 0 {
		tx = tx.Having("array_agg(article_tag.tag) @> ?", pq.Array(tags))
	}

	if categoryID != nil && *categoryID != "" {
		tx = tx.Where("category.slug = ?", *categoryID)
	}

	var articleList = make([]articleDetailWithVersion, 0)
	err := tx.Find(&articleList).Error
	if err != nil {
		return nil, e.InternalServiceError(err.Error())
	}

	var views = make([]query.ArticleView, len(articleList))
	for i, view := range articleList {
		views[i] = query.ArticleView{
			Uri:         view.ArticleDetail.URI,
			Title:       view.ArticleVersion.Title,
			Description: view.ArticleVersion.Description,
			Note:        "",
			Content:     "",
			CreatedAt:   view.FirstVersionCreatedAt.UnixMilli(),
			Category: query.ArticleCategory{
				Slug: view.Category.Slug,
				Name: view.Category.Name,
			},
			Tags: view.Tags,
		}
	}
	return views, nil
}
