package articles

import (
	"github.com/qmstar0/BlogLite-api/internal/common/domain"
	"slices"
)

type Article struct {
	domain.BaseAggregate

	uri URI

	versionList []string

	categoryID string
	tags       TagGroup

	visibility bool

	currentVersion string
}

func NewArticle(uri URI, categoryID string) *Article {
	a := &Article{
		BaseAggregate:  domain.BaseAggregate{},
		uri:            uri,
		visibility:     false,
		tags:           TagGroup{},
		categoryID:     categoryID,
		currentVersion: "",
		versionList:    []string{},
	}
	a.Emit(ArticleInitializedSuccessfullyEvent{
		URI:        a.uri.String(),
		CategoryID: categoryID,
	})
	return a
}

func (a *Article) Uri() URI {
	return a.uri
}

func (a *Article) IsVisible() bool {
	return a.visibility
}

func (a *Article) TagGroup() TagGroup {
	return a.tags
}

func (a *Article) CategoryID() string {
	return a.categoryID
}

func (a *Article) VersionList() []string {
	return slices.Clone(a.versionList)
}

func (a *Article) ChangeCategory(categoryID string) {
	if a.categoryID == categoryID {
		return
	}
	a.Emit(ArticleCategoryChangedEvent{
		URI:           a.uri.String(),
		OldCategoryID: a.categoryID,
		NewCategoryID: categoryID,
	})
	a.categoryID = categoryID
}

func (a *Article) ChangeVisibility(visibility bool) {
	if a.visibility == visibility {
		return
	}
	a.Emit(ArticleVisibilityChangedEvent{
		URI:        a.uri.String(),
		Visibility: visibility,
	})
	a.visibility = visibility
}

func (a *Article) Delete() {
	a.Emit(ArticleDeletedEvent{URI: a.uri.String()})
}

func UnmarshalArticleFromDatabase(
	uri string,
	visibility bool,
	tag []string,
	categoryID string,
	currentVersion string,
	versionList []string,
) (*Article, error) {
	return &Article{
		uri:            URI{uri},
		visibility:     visibility,
		tags:           TagGroup{tags: tag},
		categoryID:     categoryID,
		currentVersion: currentVersion,
		versionList:    versionList,
	}, nil
}
