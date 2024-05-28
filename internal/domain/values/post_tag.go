package values

import (
	"github.com/qmstar0/nightsky-api/internal/pkg/e"
	"regexp"
	"strings"
)

var TagRegexpCheck, _ = regexp.Compile(`^[\p{L}\p{N}]+$`)

type Tag string

func NewTag(s string) (Tag, error) {
	s = strings.TrimSpace(s)
	if !TagRegexpCheck.MatchString(s) {
		return "", e.DErrInvalidOperation.WithMessage("标签名格式错误")
	}
	return Tag(s), nil
}

func (t Tag) String() string {
	return string(t)
}
