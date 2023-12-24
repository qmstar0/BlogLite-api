package dto

// ArticleListViews 文章列表响应
type ArticleListViews struct {
	Paginate Paginate             `json:"paginate"`
	Items    []ArticleListDisplay `json:"items"`
}

// CateTagIndexViews 分类和标签索引响应
type CateTagIndexViews struct {
	Cate         []CateDisplay `json:"cate"`
	Tag          []TagDisplay  `json:"tags"`
	ImgUploadUrl string        `json:"imgUploadUrl"`
}

// ArticleDetailViews 文章详情响应
type ArticleDetailViews struct {
	Article      ArticleListDisplay `json:"article"`
	Cate         []CateDisplay      `json:"cate"`
	Tags         []TagDisplay       `json:"tags"`
	ImgUploadUrl string             `json:"imgUploadUrl"`
}
