package dto

// Store 需要持久化的数据模型

// ArticleStore 文化仓储模型
type ArticleStore struct {
	Title    string `json:"title" validate:"required"`
	Category int    `json:"category" validate:"required"`
	Tags     []int  `json:"tags" validate:"required"`
	Summary  string `json:"summary" validate:"required"`
	Content  string `json:"content" validate:"required"`
}

// CateStore 文章分类仓储模型
type CateStore struct {
	Name        string `form:"name" validate:"required"`
	DisplayName string `form:"displayName" validate:""`
	ParentId    int    `form:"parentId" validate:""`
	SeoDesc     string `form:"seoDesc" validate:"required"`
}

// TagStore 标签仓储模型
type TagStore struct {
	Name        string `form:"name" validate:"required"`
	DisplayName string `form:"displayName"`
	SeoDesc     string `form:"seoDesc" validate:"required"`
}

// UserStore 用户仓储模型
type UserStore struct {
	UserName string `form:"name" validate:"required"`
	//Password string
	//Email    string
}
