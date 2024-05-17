package values

import (
	"errors"
	"strings"
)

type PostUri string

func NewPostUri(uri string) (PostUri, error) {
	uri = strings.TrimSpace(uri)
	if uri == "" {
		return "", errors.New("uri不能为空")
	}
	if strings.Contains(uri, " ") {
		return "", errors.New("uri中不能存在空格")
	}
	if len([]rune(uri)) > MaxTitleCharLength {
		return "", errors.New("帖子标题过长")
	}
	return PostUri(uri), nil
}

func (p PostUri) String() string {
	return string(p)
}
