package dtoV1

// ArticleListDisplay 文章列表
type ArticleListDisplay struct {
	Article  ArticleDTO  `json:"article,omitempty"`
	Tags     []TagDTO    `json:"tags,omitempty"`
	Category CategoryDTO `json:"category,omitempty"`
	Author   UserDTO     `json:"author,omitempty"`
}

func NewArticleListDisplay(article ArticleDTO, tags []TagDTO, category CategoryDTO, author UserDTO) ArticleListDisplay {
	return ArticleListDisplay{Article: article, Tags: tags, Category: category, Author: author}
}

// ArticleListViews 文章列表和分页
type ArticleListViews struct {
	Paginate PaginateDTO          `json:"paginate"`
	Items    []ArticleListDisplay `json:"items"`
}

func NewArticleListViews(paginate PaginateDTO, items []ArticleListDisplay) ArticleListViews {
	return ArticleListViews{Paginate: paginate, Items: items}
}

// CateTagIndexViews 分类和标签索引响应
type CateTagIndexViews struct {
	Cate         []CategoryDTO `json:"cate"`
	Tag          []TagDTO      `json:"tag"`
	ImgUploadUrl string        `json:"imgUploadUrl"`
}

// ArticleDetailViews 文章详情响应
type ArticleDetailViews struct {
	Article      ArticleListDisplay `json:"article"`
	Cate         []CategoryDTO      `json:"cate"`
	Tags         []TagDTO           `json:"tags"`
	ImgUploadUrl string             `json:"imgUploadUrl"`
}
