package user

import "context"

type UserRepository interface {
	Find(ctx context.Context, uid uint32) (*User, error)
	Save(ctx context.Context, user *User) error
}
