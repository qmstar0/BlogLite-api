package dto

//数据展示模型

// ArticleListDisplay 文章详情列表
type ArticleListDisplay struct {
	Article  ArticleDisplay `json:"post,omitempty"`
	Tags     []TagDisplay   `json:"tags,omitempty"`
	Category CateDisplay    `json:"category,omitempty"`
	Author   UserDisplay    `json:"author,omitempty"`
}

// ArticleDisplay 文章内容
type ArticleDisplay struct {
	Id        int    `json:"id,omitempty"`
	Aid       string `json:"aid,omitempty"`
	Uid       string `json:"uid,omitempty"`
	TitleSlug string `json:"titleSlug,omitempty"`
	Title     string `json:"title,omitempty"`
	Summary   string `json:"summary,omitempty"`
	Content   string `json:"content,omitempty"`
	PublishAt uint   `json:"publishAt,omitempty"`
	CreateAt  uint   `json:"createAt,omitempty"`
	UpdateAt  uint   `json:"updateAt,omitempty"`
	DeleteAt  uint   `json:"deleteAt,omitempty"`
	Views     uint   `json:"views,omitempty"`
}

// TagDisplay 标签
type TagDisplay struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	SeoDesc     string `json:"seoDesc,omitempty"`
	Num         uint   `json:"num,omitempty"`
}

// CateDisplay 分类
type CateDisplay struct {
	Id          int    `json:"id,omitempty"`
	Name        string `json:"name,omitempty"`
	DisplayName string `json:"displayName,omitempty"`
	SeoDesc     string `json:"seoDesc,omitempty"`
}

// UserDisplay 作者
type UserDisplay struct {
	Uid   string `json:"uid,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Role  uint   `json:"role,omitempty"`
}

// SystemDisplay 系统
type SystemDisplay struct {
	Title        string `json:"title;omitempty"`
	Keywords     string `json:"keywords;omitempty"`
	Theme        uint   `json:"theme;omitempty"`
	Description  string `json:"description;omitempty"`
	RecordNumber string `json:"recordNumber;omitempty"`
}
