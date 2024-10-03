package query

import "context"

type ArticleDetail struct {
	URI     string
	Version *string
	Extra   bool
}

type ArticelDetailReadmodel interface {
	ArticleDetail(ctx context.Context, uri string, version *string, extra bool) (ArticleView, error)
	// ArticleDetailWithExtra 增加Note,Visibility等字段
}

type ArticleDetailhandler struct {
	rm ArticelDetailReadmodel
}

func NewArticleDetailhandler(rm ArticelDetailReadmodel) *ArticleDetailhandler {
	return &ArticleDetailhandler{rm: rm}
}

func (a *ArticleDetailhandler) Handle(ctx context.Context, query ArticleDetail) (ArticleView, error) {
	return a.rm.ArticleDetail(ctx, query.URI, query.Version, query.Extra)
}
