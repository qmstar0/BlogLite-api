package category

import "errors"

var CategoryNameFormatErr = errors.New("分类名格式错误")

type Name string

func NewName(name string) (Name, error) {
	if len(name) > 16 || len(name) <= 0 {
		return "", CategoryNameFormatErr
	}
	return Name(name), nil
}

func (n Name) String() string {
	return string(n)
}
