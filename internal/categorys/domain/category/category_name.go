package category

import (
	"common/e"
	"common/idtools"
	"errors"
	"strings"
)

var CategoryNameFormatErr = errors.New("分类名格式错误")

type Name string

func NewName(name string) (Name, error) {
	if len(name) > 16 || len(name) <= 0 {
		return "", e.Wrap(e.ValueObjectCheckErr, CategoryNameFormatErr)
	}
	return Name(strings.TrimSpace(name)), nil
}

func (n Name) String() string {
	return string(n)
}

func (n Name) ToUint32() uint32 {
	return idtools.NewHashID(string(n))
}
