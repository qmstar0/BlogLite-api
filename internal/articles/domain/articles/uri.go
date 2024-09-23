package articles

import (
	"github.com/qmstar0/BlogLite-api/internal/common/e"
	"regexp"
)

var uriFormatRe = regexp.MustCompile("^[a-zA-Z0-9_-]+$")

type URI struct {
	s string
}

func NewUri(s string) URI {
	return URI{s: s}
}

func (u URI) CheckFormat() error {
	if !uriFormatRe.MatchString(u.s) {
		return e.InvalidActionError("uri格式错误，uri只能包含字母、数字、下划线或连字符")
	}
	return nil
}

func (u URI) String() string {
	return u.s
}
