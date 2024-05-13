package domain

import (
	"github.com/uptrace/bun"
	"go-blog-ddd/internal/application/query"
	"go-blog-ddd/internal/domain/aggregates"
	"go-blog-ddd/internal/domain/values"
	"time"
)

type PostM struct {
	bun.BaseModel `bun:"table:post,alias:post"`
	ID            uint32 `bun:"id,pk"`

	Uri   string
	Title string
	Desc  string

	Content string

	CategoryID uint32     `bun:"category_id"`
	Category   *CategoryM `bun:"rel:has-one,join:category_id=id"`

	Visible  bool        `bun:"visible"`
	PostTags []*PostTagM `bun:"rel:has-many,join:id=post_id"`

	CreatedAt time.Time `bun:"created_at"`
	UpdatedAt time.Time `bun:"updated_at"`
}

func PostAggregateToModel(post *aggregates.Post) *PostM {
	var tags = make([]*PostTagM, len(post.Tags))
	for i, tag := range post.Tags {
		tags[i] = &PostTagM{
			PostID: post.ID,
			Tag:    tag.String(),
		}
	}
	return &PostM{
		ID:         post.ID,
		Uri:        post.Uri.String(),
		Title:      post.Title.String(),
		Desc:       post.Desc,
		Content:    post.Content,
		Visible:    post.Visible,
		CategoryID: post.CategoryID,
		PostTags:   tags,
		CreatedAt:  post.CreatedAt,
		UpdatedAt:  post.UpdatedAt,
	}
}

func PostModelToAggregate(post *PostM) *aggregates.Post {
	var tags = make([]values.Tag, len(post.PostTags))
	for i, tag := range post.PostTags {
		tags[i] = values.Tag(tag.Tag)
	}
	return &aggregates.Post{
		ID:         post.ID,
		Content:    post.Content,
		Uri:        values.PostUri(post.Uri),
		Title:      values.PostTitle(post.Title),
		Desc:       post.Desc,
		Tags:       tags,
		Visible:    post.Visible,
		CategoryID: post.CategoryID,
		CreatedAt:  post.CreatedAt,
		UpdatedAt:  post.UpdatedAt,
	}
}

func PostModelToView(m *PostM) query.PostView {
	var cate *query.CategoryView
	if m.Category != nil {
		view := CategoryModelToView(m.Category)
		cate = &view
	}

	var tags = make([]string, len(m.PostTags))
	for i, tag := range m.PostTags {
		tags[i] = tag.Tag
	}

	return query.PostView{
		ID:        m.ID,
		Uri:       m.Uri,
		Title:     m.Title,
		Content:   m.Content,
		Desc:      m.Desc,
		Tags:      tags,
		Visible:   m.Visible,
		Category:  cate,
		CreatedAt: m.CreatedAt.UnixMilli(),
		UpdatedAt: m.UpdatedAt.UnixMilli(),
	}
}

func PostModelToListView(postM []*PostM) query.PostListView {
	postMLen := len(postM)
	result := query.PostListView{
		Count: postMLen,
		Page:  1,
		Items: make([]query.PostView, postMLen),
	}
	if postMLen <= 0 {
		return result
	}
	for i, m := range postM {
		result.Items[i] = PostModelToView(m)
	}
	return result
}
