package domain

import (
	"github.com/uptrace/bun"
	"go-blog-ddd/internal/application/query"
	"go-blog-ddd/internal/domain/aggregates"
	"go-blog-ddd/internal/domain/values"
	"time"
)

type PostTagM struct {
	bun.BaseModel `bun:"table:post_tag,alias:post_tag"`
	ID            uint32 `bun:"id,pk,autoincrement"`
	PostID        uint32 `bun:"post_id"`
	Tag           string `bun:"tag"`
}

type PostM struct {
	bun.BaseModel `bun:"table:post,alias:post"`
	ID            uint32 `bun:"id,pk"`

	Uri   string
	Title string
	Desc  string

	SourceFilePath string `bun:"source_file_path"`
	Content        string

	CategoryID uint32     `bun:"category_id"`
	Category   *CategoryM `bun:"rel:has-one,join:category_id=id"`

	Visible  bool
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
		ID:             post.ID,
		Uri:            post.Uri.String(),
		Title:          post.Title.String(),
		Desc:           post.Desc,
		SourceFilePath: post.SourcePath,
		Content:        post.Content,
		Visible:        post.Visible,
		CategoryID:     post.CategoryID,
		PostTags:       tags,
		CreatedAt:      post.CreatedAt,
		UpdatedAt:      post.UpdatedAt,
	}
}

func PostModelToAggregate(post *PostM) *aggregates.Post {
	var tags = make([]values.Tag, len(post.PostTags))
	for i, tag := range post.PostTags {
		tags[i] = values.Tag(tag.Tag)
	}
	return &aggregates.Post{
		ID:         post.ID,
		SourcePath: post.SourceFilePath,
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
	var cateview *query.CategoryView
	if m.Category != nil {
		view := CategoryModelToView(m.Category)
		cateview = &view
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
		Visible:   m.Visible,
		Tags:      tags,
		Category:  cateview,
		CreatedAt: m.CreatedAt.UnixMilli(),
		UpdatedAt: m.UpdatedAt.UnixMilli(),
	}
}
