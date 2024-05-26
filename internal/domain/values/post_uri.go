package values

import (
	"fmt"
	"github.com/qmstar0/domain/internal/pkg/e"

	"strings"
)

const MaxUriCharLength = 75

type PostUri string

func NewPostUri(uri string) (PostUri, error) {
	uri = strings.TrimSpace(uri)
	if uri == "" || strings.Contains(uri, " ") {
		return "", e.DErrInvalidOperation.WithMessage("uri格式错误")
	}
	if len([]rune(uri)) > MaxUriCharLength {
		return "", e.DErrInvalidOperation.WithMessage(fmt.Sprintf("帖子标题过长(最多%d)", MaxUriCharLength))
	}
	return PostUri(uri), nil
}

func (p PostUri) String() string {
	return string(p)
}
