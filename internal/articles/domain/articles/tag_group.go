package articles

import (
	"fmt"
	"github.com/qmstar0/BlogLite-api/internal/common/constant"
	"github.com/qmstar0/BlogLite-api/internal/common/e"
	"sort"
	"strings"
)

type TagGroup struct {
	tags []string
}

func NewTagGroup(tags []string) (TagGroup, error) {
	tagSet := make(map[string]struct{})
	for _, tag := range tags {
		tagSet[strings.TrimSpace(tag)] = struct{}{}
	}

	if len(tagSet) > constant.ArticleTagsMaxNum {
		return TagGroup{}, e.InvalidActionError(fmt.Sprintf("一个文章最多只能有%d个标签", constant.ArticleTagsMaxNum))
	}
	result := make([]string, 0)
	for s, _ := range tagSet {
		result = append(result, s)
	}
	sort.Strings(result)
	return TagGroup{tags: result}, nil
}

func (t TagGroup) Value() []string {
	value := make([]string, len(t.tags))
	copy(value, t.tags)
	return value
}

func (a *Article) ChangeTagGroup(tagGroup TagGroup) {
	a.Emit(ArticleTagsModifiedEvent{
		URI:     a.uri.String(),
		OldTags: a.tags.Value(),
		NewTags: tagGroup.Value(),
	})
	a.tags = tagGroup
}
