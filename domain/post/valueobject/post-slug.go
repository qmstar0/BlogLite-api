package valueobject

import (
	"errors"
	"strings"
)

type PostSlug string

var SlugFormatErr = errors.New("post-content slug format error")

func NewPostSlug(slug string) (PostSlug, error) {
	if strings.Contains(slug, " ") {
		return "", SlugFormatErr
	}
	return PostSlug(slug), nil
}

func (p PostSlug) ToString() string {
	return string(p)
}
