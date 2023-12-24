package dto

//索引页

// IndexArticleList 文章列表
type IndexArticleList struct {
	Paginate Paginate
	Items    []*ArticleListDisplay
}

// IndexArticle 文章内容
type IndexArticle struct {
	Id        int    `json:"id,omitempty"`
	Aid       string `json:"aid,omitempty"`
	Uid       string `json:"uid,omitempty"`
	Title     string `json:"title,omitempty"`
	Summary   string `json:"summary,omitempty"`
	Content   string `json:"content,omitempty"`
	PublishAt uint   `json:"deletedAt,omitempty"`
	CreatedAt uint   `json:"createdAt,omitempty"`
	UpdatedAt uint   `json:"updatedAt,omitempty"`
	Views     uint   `json:"views,omitempty"`
}

type IndexRss struct {
	Id        int    `json:"id,omitempty"`
	Uid       string `json:"uid,omitempty"`
	Author    string `json:"author,omitempty"`
	Title     string `json:"title,omitempty"`
	Summary   string `json:"summary,omitempty"`
	CreatedAt uint   `json:"createdAt,omitempty"`
	UpdatedAt uint   `json:"updatedAt,omitempty"`
}

// IndexArticleDetail 文章详情索引
type IndexArticleDetail struct {
	Post     IndexArticle `json:"posts-content,omitempty"`
	Tags     []TagDisplay `json:"tags,omitempty"`
	Category CateDisplay  `json:"categorys,omitempty"`
	Author   UserDisplay  `json:"author,omitempty"`
}
