package aggregates

import (
	"context"
	"errors"
	"github.com/qmstar0/domain/internal/domain/values"
	"github.com/qmstar0/domain/internal/pkg/e"
	"time"
)

type Post struct {
	AggregateRoot

	ID uint32

	Content string

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
	content string,
) *Post {
	now := time.Now()
	return &Post{
		ID:        pid,
		Uri:       uri,
		Content:   content,
		CreatedAt: now,
		UpdatedAt: now,
		Visible:   false,
	}
}

func (p *Post) ModifyVisible(visible bool) error {
	if visible {
		if p.Title == "" || p.Desc == "" {
			return e.DErrInvalidOperation.WithMessage("请先完善帖子的标题和简介")
		}
	}
	p.Visible = visible
	p.UpdatedAt = time.Now()
	return nil
}

func (p *Post) ModifyPostTags(tags []values.Tag) error {
	if len(tags) > 4 {
		return e.DErrInvalidOperation.WithError(errors.New("Post最多拥有4个Tag")).WithMessage("每个帖子最多只能设置4个标签")
	}
	p.Tags = tags
	return nil
}

func (p *Post) ModifyPostTitle(title values.PostTitle) {
	p.Title = title
	p.UpdatedAt = time.Now()
}

func (p *Post) ResetContent(content string) {
	p.Content = content
	p.UpdatedAt = time.Now()
}

func (p *Post) ModifyPostDesc(desc string) {
	p.Desc = desc
	p.UpdatedAt = time.Now()
}

func (p *Post) Delete() {}

type PostRepository interface {
	FindByUri(ctx context.Context, uri values.PostUri) (*Post, error)
	FindOrErrByUri(ctx context.Context, uri values.PostUri) (*Post, error)
	FindOrErrByID(ctx context.Context, id uint32) (*Post, error)
	Save(ctx context.Context, post *Post) error
	NextID(ctx context.Context) (uint32, error)
	Delete(ctx context.Context, pid uint32) error
	ResourceUniquenessCheck(ctx context.Context, uri values.PostUri) error
}
