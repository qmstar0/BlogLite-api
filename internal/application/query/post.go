package query

import (
	"context"
)

type PostView struct {
	ID      uint32 `json:"id,omitempty"`
	Uri     string `json:"uri,omitempty"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Desc    string `json:"desc,omitempty"`

	Tags []string `json:"tags,omitempty"`

	Visible bool `json:"visible,omitempty"`

	Category  *CategoryView `json:"category,omitempty"`
	CreatedAt int64         `json:"createdAt,omitempty"`
	UpdatedAt int64         `json:"updatedAt,omitempty"`
}

type PostListView struct {
	Count int        `json:"count"`
	Page  int        `json:"page,omitempty"`
	Items []PostView `json:"items"`
}

type PostQueryControl interface {
	FindByID(ctx context.Context, pid uint32) (PostView, error)
	FindByUri(ctx context.Context, uri string) (PostView, error)
	RecentPosts(ctx context.Context, limit int) (PostListView, error)
	AllWithFilter(
		ctx context.Context,
		limit,
		offset int,
		tags []string,
		categroyID uint32,
		onlyVisible bool,
	) (PostListView, error)
}
