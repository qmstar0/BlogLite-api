package users

import (
	"context"
)

type RepoUser interface {
	NewUser(c context.Context, user *User) error
	UptUser(c context.Context, user *User) error
	DelUser(c context.Context, user *User) error
	GetUser(c context.Context, user *User) (*User, error)
}
