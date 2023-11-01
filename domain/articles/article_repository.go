package articles

import (
	"context"
)

// RepoArticleMate 文章仓储
type RepoArticleMate interface {
	NewArticle(c context.Context, art *ArticleMate) error
	UptArticle(c context.Context, art *ArticleMate) error
	DelArticle(c context.Context, art *ArticleMate) error
	GetArticle(c context.Context, art *ArticleMate) (*ArticleMate, error)
	AllArticle(c context.Context, limit int, offset int, isDraft, isTrash bool) ([]*ArticleMate, error)
}

// RepoTags 文章标签仓储
type RepoTags interface {
	NewTag(c context.Context, tags *ArticleTags) error
	UptTag(c context.Context, tags *ArticleTags) error
	DelTag(c context.Context, tags *ArticleTags) error
	GetTag(c context.Context, tags []int) ([]*ArticleTags, error)
	AllTag(c context.Context) ([]*ArticleTags, error)
}

// RepoCate 文章分类仓储
type RepoCate interface {
	NewCate(c context.Context, cate *ArticleCategory) error
	UptCate(c context.Context, cate *ArticleCategory) error
	DelCate(c context.Context, cate *ArticleCategory) error
	GetCate(c context.Context, cate *ArticleCategory) (*ArticleCategory, error)
	AllCate(c context.Context) ([]*ArticleCategory, error)
}
