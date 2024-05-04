package aggregates

import (
	"context"
	"fmt"
	"go-blog-ddd/internal/domain/values"
	"strings"
	"time"
)

type Post struct {
	AggregateRoot

	ID uint32

	SourcePath string
	Content    string

	Uri   values.PostUri
	Title values.PostTitle
	Desc  string

	Tags []values.Tag

	Visible    bool
	CategoryID uint32

	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewPost(
	pid uint32,
	uri values.PostUri,
	sourcePath,
	content string,
) *Post {
	return &Post{
		ID:         pid,
		Uri:        uri,
		Content:    content,
		SourcePath: sourcePath,
		CreatedAt:  time.Now(),
		Visible:    false,
	}
}

func (p *Post) ModifyVisible(visible bool) error {
	if visible {
		var params = make([]string, 0)
		if p.Title == "" {
			params = append(params, "title(Post标题)")
		}
		if p.Desc == "" {
			params = append(params, "desc(Post简介)")
		}
		if len(params) > 0 {
			return fmt.Errorf("你必须设置%s等参数后，才能设置Post为可见", strings.Join(params, "; "))
		}
	}
	p.Visible = visible
	p.UpdatedAt = time.Now()
	return nil
}

func (p *Post) ModifyPostTags(tags []values.Tag) {
	p.Tags = tags
}

func (p *Post) ModifyPostInfo(title values.PostTitle, desc string) {
	p.Desc = desc
	p.Title = title
	p.UpdatedAt = time.Now()
}

func (p *Post) ResetContent(sourcePath, content string) {
	p.UpdatedAt = time.Now()
	p.Content = content
	p.SourcePath = sourcePath
}

func (p *Post) Delete() {}

type PostRepository interface {
	//FindByTags(ctx context.Context, tags []values.Tag) ([]*Post, error)
	FindByUri(ctx context.Context, uri values.PostUri) (*Post, error)
	FindOrErrByUri(ctx context.Context, uri values.PostUri) (*Post, error)
	FindOrErrByID(ctx context.Context, id uint32) (*Post, error)
	Save(ctx context.Context, post *Post) error
	NextID(ctx context.Context) uint32
	Delete(ctx context.Context, pid uint32) error
}
