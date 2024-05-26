package values

import (
	"github.com/qmstar0/domain/internal/pkg/e"
	"regexp"
	"strings"
)

var CategoryRegexpCheck, _ = regexp.Compile(`^[\p{L}\p{N}]+$`)

type CategoryName string

func NewCategoryName(s string) (CategoryName, error) {
	s = strings.TrimSpace(s)
	if !CategoryRegexpCheck.MatchString(s) {
		return "", e.DErrInvalidOperation.WithMessage("分类名格式错误")
	}
	return CategoryName(s), nil
}

func (n CategoryName) String() string {
	return string(n)
}
