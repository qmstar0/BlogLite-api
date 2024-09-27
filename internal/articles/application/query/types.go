package query

type ArticleView struct {
	Uri         string          `json:"uri"`
	Title       string          `json:"title"`
	Description string          `json:"description"`
	Note        string          `json:"note,omitempty"`
	Content     string          `json:"content,omitempty"`
	Visibility  bool            `json:"visibility"`
	CreatedAt   int64           `json:"createdAt"`
	Category    ArticleCategory `json:"category"`
	Tags        []string        `json:"tags"`
}

type ArticleCategory struct {
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
