package values

import (
	"errors"
	"strings"
)

const MaxTitleCharLength = 50

type PostTitle string

func NewPostTitle(title string) (PostTitle, error) {
	title = strings.TrimSpace(title)
	if title == "" {
		return "", errors.New("分类名不能为空")
	}
	if len([]rune(title)) > MaxTitleCharLength {
		return "", errors.New("帖子标题过长")
	}
	return PostTitle(title), nil
}

func (p PostTitle) String() string {
	return string(p)
}
