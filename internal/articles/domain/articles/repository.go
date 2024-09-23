package articles

import "context"

type ArticleRepository interface {
	Find(ctx context.Context, uri URI) (*Article, error)
	Save(ctx context.Context, article *Article) error
	UpdateArticle(ctx context.Context, uri URI, updateFn func(*Article) (*Article, error)) error
	Remove(ctx context.Context, article *Article) error
}
