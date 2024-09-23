package adapter

import (
	"context"
	"github.com/qmstar0/BlogLite-api/internal/common/e"
	"gorm.io/gorm"
)

type ArticleVersionDuplicationCheckService struct {
	db *gorm.DB
}

func NewArticleVersionDuplicationCheckService(db *gorm.DB) *ArticleVersionDuplicationCheckService {
	return &ArticleVersionDuplicationCheckService{db: db}
}

func (a ArticleVersionDuplicationCheckService) CheckDuplication(ctx context.Context, versionHash string) error {
	var count int64
	err := a.db.WithContext(ctx).
		Model(&ArticleVersion{}).
		Where("version = ?", versionHash).
		Limit(1).
		Count(&count).Error
	if err != nil {
		return e.InternalServiceError(err.Error())
	}
	if count != 0 {
		return e.InvalidActionError("该版本已存在或已上传至其他文章")
	}

	return nil
}
