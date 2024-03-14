package post

import "common/idtools"

type Uri string

func NewUri(s string) (Uri, error) {
	return Uri(s), nil
}

func (u Uri) ToID() uint32 {

	return idtools.NewHashID([]byte(u))
}

func (u Uri) String() string {
	return string(u)
}
