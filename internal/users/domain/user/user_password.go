package user

import (
	"common/e"
	"errors"
)

type Password string

func (p Password) String() string {
	return string(p)
}

func NewPassword(pwd string) (Password, error) {
	if pwd == "" {
		return "", e.Wrap(e.PasswordFormatErr, errors.New("password format err"))
	}
	return Password(pwd), nil
}
