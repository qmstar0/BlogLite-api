package query

import "context"

type ArticleMetadata struct {
	URI string
}

type ArticelMetadataReadmodel interface {
	ArticleMetadata(ctx context.Context, uri string) (ArticleMetadataView, error)
}

type ArticleMetadatahandler struct {
	rm ArticelMetadataReadmodel
}

func NewArticleMetadatahandler(rm ArticelMetadataReadmodel) *ArticleMetadatahandler {
	return &ArticleMetadatahandler{rm: rm}
}

func (a *ArticleMetadatahandler) Handle(ctx context.Context, query ArticleMetadata) (ArticleMetadataView, error) {
	return a.rm.ArticleMetadata(ctx, query.URI)
}
