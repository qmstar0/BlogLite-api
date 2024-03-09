package user

import "common/idtools"

type Email string

func (m Email) String() string {
	return string(m)
}

func (m Email) ToID() uint32 {
	return idtools.NewHashID(string(m))
}

func NewEmail(email string) (Email, error) {
	return Email(email), nil
}
