package dto

import (
	"context"
)

type DTO interface {
	ArticleR
	ArticleW
	CateR
	CateW
	TagsR
	TagsW
}

// ArticleR 文章读取接口
type ArticleR interface {
	GetArticleDetailList(c context.Context, limit int, offset int, status uint) ([]ArticleListDisplay, error)
	GetArticleDetail(c context.Context, aid string) (ArticleListDisplay, error)
	GetArticle(c context.Context, aid string) (ArticleDisplay, error)
}

// TagsR 标签读取接口
type TagsR interface {
	GetTag(c context.Context, tagId []int) ([]TagDisplay, error)
	AllTag(c context.Context) ([]TagDisplay, error)
}

// CateR 文章分类读取接口
type CateR interface {
	GetCate(c context.Context, cateId int) (CateDisplay, error)
	GetCateByParentId(c context.Context, ParentId int) (CateDisplay, error)
	AllCate(c context.Context) ([]CateDisplay, error)
}

// ArticleW 文章写入接口
type ArticleW interface {
	NewArticle(c context.Context, uid string, store ArticleStore) error
	UpdateArticle(c context.Context, aid string, store ArticleStore) error
	DeleteArtcle(c context.Context, aid string) error
	PublishArticle(c context.Context, aid string) error
}

// TagsW 标签写入接口
type TagsW interface {
	NewTag(c context.Context, store TagStore) error
	UpdateTag(c context.Context, tagId int, store TagStore) error
	DeleteTag(c context.Context, tagId int) error
}

// CateW 文章分类写入接口
type CateW interface {
	NewCate(c context.Context, store CateStore) error
	UpdateCate(c context.Context, cateId int, store CateStore) error
	DeleteCate(c context.Context, cateId int) error
}
