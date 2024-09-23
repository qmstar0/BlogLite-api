package query

import "context"

type ArticleVersionList struct {
	Uri string
}

type ArticleVersionListReadmodel interface {
	ArticleVersionList(ctx context.Context, uri string) ([]ArticleVersionView, error)
}

type ArticleVersionListHandler struct {
	rm ArticleVersionListReadmodel
}

func NewArticleVersionListHandler(rm ArticleVersionListReadmodel) *ArticleVersionListHandler {
	return &ArticleVersionListHandler{rm: rm}
}

func (a *ArticleVersionListHandler) Handle(ctx context.Context, query ArticleVersionList) (ArticleVersionListView, error) {
	list, err := a.rm.ArticleVersionList(ctx, query.Uri)
	if err != nil {
		return ArticleVersionListView{}, err
	}
	return ArticleVersionListView{
		Count: len(list),
		Items: list,
	}, nil
}
