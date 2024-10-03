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
	Visibility            bool
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
		Visibility:     false,
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
		Where("uri = ?", event.URI).Update("visibility", event.Visibility).Error
	if err != nil {
		return err
	}
	return nil
}

func (p PostgresArticleDetailReadmodel) joinsTable(db *gorm.DB) *gorm.DB {
	return db.Joins("JOIN article_version ON article_version.uri = article_detail.uri").
		Joins("LEFT JOIN article_tag ON article_detail.uri = article_tag.article_uri").
		Joins("LEFT JOIN category ON article_detail.category_id = category.slug").
		Group("article_detail.uri, category.slug, article_version.version")
}

func (p PostgresArticleDetailReadmodel) ArticleDetail(ctx context.Context, uri string, version *string, extra bool) (query.ArticleView, error) {
	queryFields := defaultSelectFields + fieldsArticleVersionContent

	if extra {
		queryFields += extraFields
	}

	tx := p.joinsTable(p.db.WithContext(ctx).
		Model(&ArticleDetail{}).
		Select(queryFields)).
		Limit(1).
		Where("article_detail.uri = ?", uri).
		Where("article_detail.visibility = true")

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
		Title:       article.ArticleVersion.Title,
		Version:     article.ArticleDetail.CurrentVersion,
		Description: article.ArticleVersion.Description,
		Note:        article.ArticleVersion.Note,
		Content:     article.ArticleVersion.Content,
		Visibility:  article.ArticleDetail.Visibility,
		CreatedAt:   article.ArticleDetail.FirstVersionCreatedAt.UnixMilli(),
		Category: query.ArticleCategory{
			Slug: article.Category.Slug,
			Name: article.Category.Name,
		},
		Tags: article.Tags,
	}, nil
}

func (p PostgresArticleDetailReadmodel) ArticleList(
	ctx context.Context,
	offset, limit int,
	tags []string,
	categoryID *string,
	extra bool,
) ([]query.ArticleView, error) {

	var queryFields = defaultSelectFields
	if extra {
		queryFields += extraFields
	}
	tx := p.joinsTable(p.db.WithContext(ctx).
		Model(&ArticleDetail{}).
		Select(string(queryFields))).
		Offset(offset).
		Limit(limit).
		Where("article_version.version = article_detail.current_version")
	if len(tags) != 0 {
		tx = tx.Having("array_agg(article_tag.tag) @> ?", pq.Array(tags))
	}
	tx = tx.Omit("article_version.content, article_version.note")

	if categoryID != nil && *categoryID != "" {
		tx = tx.Where("category.slug = ?", *categoryID)
	}

	if !extra {
		tx = tx.Where("article_detail.visibility = true")
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
			Version:     view.ArticleDetail.CurrentVersion,
			Description: view.ArticleVersion.Description,
			Note:        view.ArticleVersion.Note,
			Content:     view.ArticleVersion.Content,
			Visibility:  view.ArticleDetail.Visibility,
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

const defaultSelectFields = "article_detail.uri, " +
	"article_detail.first_version_created_at, " +
	"article_detail.current_version, " +
	"category.slug, " +
	"category.name, " +
	"article_version.title, " +
	"article_version.description, " +
	"array_remove(array_agg(article_tag.tag), null) AS tags"
const fieldsArticleVersionContent = ", article_version.content"
const extraFields = ", article_version.note, article_detail.visibility"
