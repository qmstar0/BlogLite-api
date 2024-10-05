package query

type ArticleView struct {
	Uri         string              `json:"uri"`
	Title       string              `json:"title"`
	Version     string              `json:"version,omitempty"`
	Description string              `json:"description"`
	Note        string              `json:"note,omitempty"`
	Content     string              `json:"content,omitempty"`
	Visibility  bool                `json:"visibility,omitempty"`
	CreatedAt   int64               `json:"createdAt"`
	Category    ArticleCategoryView `json:"category"`
	Tags        []string            `json:"tags"`
}

type ArticleCategoryView struct {
	Slug string `json:"slug"`
	Name string `json:"name"`
}

type ArticleListView struct {
	Count int           `json:"count"`
	Page  int           `json:"page"`
	Items []ArticleView `json:"items"`
	Prev  bool          `json:"prev"`
	Next  bool          `json:"next"`
}

type ArticleVersionView struct {
	Version   string `json:"version"`
	Note      string `json:"note"`
	Title     string `json:"title"`
	CreatedAt int64  `json:"createdAt"`
}

type ArticleVersionListView struct {
	Count int                  `json:"count"`
	Items []ArticleVersionView `json:"items"`
}

type TagListView struct {
	Count int      `json:"count"`
	Items []string `json:"items"`
}

type ArticleMetadataView struct {
	URI                   string              `json:"uri"`
	Version               string              `json:"version"`
	Visibility            bool                `json:"visibility"`
	Category              ArticleCategoryView `json:"category"`
	FirstVersionCreatedAt int64               `json:"firstVersionCreatedAt"`
	Tags                  []string            `json:"tags"`
}

type ArticleMetadataListView struct {
	Count int                   `json:"count"`
	Page  int                   `json:"page"`
	Items []ArticleMetadataView `json:"items"`
	Prev  bool                  `json:"prev"`
	Next  bool                  `json:"next"`
}
