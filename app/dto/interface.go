package dto

import (
	"blog/domain/users/token"
	"context"
)

// Authorizer 验证器
type Authorizer interface {
	GenCaptchaToken(email, captcha string) (string, error)
	SendCaptchaEmail(email, captcha string) error
	VaildateCaptcha(cap Captcha) error
	GenAuthToken(user UserDisplay) (string, error)
	VaildateAuth(authToken string) (*token.AuthClaims, error)
}

// ArticleR 文章读取接口
type ArticleR interface {
	GetArticleDetailList(c context.Context, limit int, offset int, isDraft bool, isTrash bool) ([]ArticleListDisplay, error)
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

// UserR 用户数据读取接口
type UserR interface {
	GetUser(c context.Context, email string) (UserDisplay, error)
	GetUserByUid(c context.Context, uid string) (UserDisplay, error)
}

// ArticleW 文章写入接口
type ArticleW interface {
	NewArticle(c context.Context, uid string, store ArticleStore) error
	UpdateArticle(c context.Context, aid string, store ArticleStore) error
	DeleteArtcle(c context.Context, aid string) error
	PublishArticle(c context.Context, aid string) error
}

// UserW 用户写入接口
type UserW interface {
	NewUser(c context.Context, email string) (UserDisplay, error)
	UpdateUser(c context.Context, email string, store UserStore) error
	ResetPwd(c context.Context, email, newPwd string) error
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
