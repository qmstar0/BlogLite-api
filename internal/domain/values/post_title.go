package values

import (
	"fmt"
	"github.com/qmstar0/nightsky-api/internal/pkg/e"

	"strings"
)

const MaxTitleCharLength = 50

type PostTitle string

func NewPostTitle(title string) (PostTitle, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return "", e.DErrInvalidOperation.WithMessage("分类名不能为空")
	}
	if len([]rune(title)) > MaxTitleCharLength {
		return "", e.DErrInvalidOperation.WithMessage(fmt.Sprintf("帖子标题过长(最多%d个字符)", MaxTitleCharLength))
	}
	return PostTitle(title), nil
}

func (p PostTitle) String() string {
	return string(p)
}
