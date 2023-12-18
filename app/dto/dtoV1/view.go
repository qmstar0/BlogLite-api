package dtoV1

import (
	"blog/domain/articles"
	"blog/domain/users"
	"blog/infra/paginate"
)

// ArticleDTO 文章内容
type ArticleDTO struct {
	Id        int    `json:"id"`
	Aid       string `json:"aid,omitempty"`
	Uid       string `json:"uid,omitempty"`
	TitleSlug string `json:"titleSlug,omitempty"`
	Title     string `json:"title,omitempty"`
	Summary   string `json:"summary,omitempty"`
	Content   string `json:"content,omitempty"`
	PublishAt uint   `json:"publishAt"`
	CreateAt  uint   `json:"createAt"`
	UpdateAt  uint   `json:"updateAt"`
	DeleteAt  uint   `json:"deleteAt"`
	Views     uint   `json:"views"`
}

func NewArticleDTO(article *articles.ArticleMate) ArticleDTO {
	return ArticleDTO{}
}

// TagDTO 标签
type TagDTO struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	SeoDesc     string `json:"seoDesc,omitempty"`
	Num         uint   `json:"num,omitempty"`
}

func NewTagDTO(tag *articles.ArticleTags) TagDTO {
	return TagDTO{}
}

// CategoryDTO 分类
type CategoryDTO struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	SeoDesc     string `json:"seoDesc,omitempty"`
}

func NewCategoryDTO(cate *articles.ArticleCategory) CategoryDTO {
	return CategoryDTO{}
}

// UserDTO 作者
type UserDTO struct {
	Uid   string `json:"uid,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Role  uint   `json:"role,omitempty"`
}

func NewUserDTO(user *users.User) UserDTO {
	return UserDTO{}
}

// PaginateDTO 分页器
type PaginateDTO struct {
	Limit   int `json:"limit"`
	Count   int `json:"count"`
	Total   int `json:"total"`
	Last    int `json:"last"`
	Current int `json:"current"`
	Next    int `json:"next"`
}

func NewPaginateDTO(paginate *paginate.Paginate) PaginateDTO {
	return PaginateDTO{}
}

// SystemDTO 系统
type SystemDTO struct {
	Title        string `json:"title;omitempty"`
	Keywords     string `json:"keywords;omitempty"`
	Theme        uint   `json:"theme;omitempty"`
	Description  string `json:"description;omitempty"`
	RecordNumber string `json:"recordNumber;omitempty"`
}

func NewSystemDTO() SystemDTO {
	return SystemDTO{}
}
