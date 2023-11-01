package dto

// ArticleListViews 文章列表响应
type ArticleListViews struct {
	Paginate Paginate
	Items    []ArticleListDisplay
}

// CateTagIndexViews 分类和标签索引响应
type CateTagIndexViews struct {
	Cate         []CateDisplay
	Tag          []TagDisplay
	ImgUploadUrl string
}

// ArticleDetailViews 文章详情响应
type ArticleDetailViews struct {
	Article      ArticleListDisplay
	Cate         []CateDisplay
	Tags         []TagDisplay
	ImgUploadUrl string
}
