package category

import "common/idtools"

type Name string

func NewName(s string) (Name, error) {
	return Name(s), nil
}
func (n Name) ToID() uint32 {
	return idtools.NewHashID([]byte(n))
}

func (n Name) String() string {
	return string(n)
}
