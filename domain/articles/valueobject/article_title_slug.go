package valueobject

import (
	"blog/infra/e"
	"strings"
)

type TitleSlug string

func NewTitleSlug(slug string) (TitleSlug, error) {
	if strings.Contains(slug, " ") {
		return "", e.NewError(e.InvalidParam, nil)
	}
	return TitleSlug(slug), nil
}

func (t *TitleSlug) ToString() string {
	return string(*t)
}
