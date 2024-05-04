package query

import (
	"context"
	"go-blog-ddd/internal/adapter/utils"
)

type PostView struct {
	ID      uint32 `json:"id,omitempty"`
	Uri     string `json:"uri,omitempty"`
	Title   string `json:"title,omitempty"`
	Content string `json:"content,omitempty"`
	Desc    string `json:"desc,omitempty"`
	Visible bool   `json:"visible,omitempty"`

	Tags []string `json:"tags,omitempty"`

	Category  *CategoryView `json:"category,omitempty"`
	CreatedAt int64         `json:"createdAt,omitempty"`
	UpdatedAt int64         `json:"updatedAt,omitempty"`
}

type PostListView struct {
	Count int        `json:"count"`
	Page  int        `json:"page"`
	Items []PostView `json:"items"`
}
type PostReadModel interface {
	FindByID(ctx context.Context, pid uint32) (PostView, error)
	AllWithFilter(ctx context.Context, limit, offset int, tags []string, categroyID uint32) ([]PostView, error)
}

type PostQueryControl struct {
	readmodel PostReadModel
}

func NewPostQueryControl(readmodel PostReadModel) PostQueryControl {
	return PostQueryControl{readmodel: readmodel}
}

func (p PostQueryControl) FindByID(ctx context.Context, pid uint32) (PostView, error) {
	return p.readmodel.FindByID(ctx, pid)
}

func (p PostQueryControl) GetPostsWithFilter(
	ctx context.Context,
	limit, page int,
	tags []string,
	categroyID uint32,
) (PostListView, error) {
	offset := utils.Offset(page, limit)

	postViews, err := p.readmodel.AllWithFilter(ctx, limit, offset, tags, categroyID)
	if err != nil {
		return PostListView{}, err
	}
	return PostListView{
		Count: len(postViews),
		Page:  page,
		Items: postViews,
	}, nil
}
