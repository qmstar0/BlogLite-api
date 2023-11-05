package dto

// Store 需要持久化的数据模型

// ArticleStore 文化仓储模型
type ArticleStore struct {
	Title     string `json:"title" validate:"required"`
	TitleSlug string `json:"titleSlug" validate:"required"`
	Category  int    `json:"category" validate:"required"`
	Tags      []int  `json:"tags" validate:"required"`
	Summary   string `json:"summary" validate:"required"`
	Content   string `json:"content" validate:"required"`
}

// CateStore 文章分类仓储模型
type CateStore struct {
	Name        string `json:"name" form:"name" validate:"required"`
	DisplayName string `json:"display" form:"display"`
	ParentId    int    `json:"parentId" form:"parentId"`
	SeoDesc     string `json:"seoDesc" form:"seoDesc" validate:"required"`
}

// TagStore 标签仓储模型
type TagStore struct {
	Name        string `json:"name" form:"name" validate:"required"`
	DisplayName string `json:"display" form:"display"`
	SeoDesc     string `json:"seoDesc" form:"seoDesc" validate:"required"`
}

// UserStore 用户仓储模型
type UserStore struct {
	UserName string `json:"name" form:"name" validate:"required"`
	//Password string
	//Email    string
}
