package user

import "context"

type UserRepository interface {
	Exist(ctx context.Context, uid uint32) (bool, error)
	Find(ctx context.Context, uid uint32) (*User, error)
	Save(ctx context.Context, user *User) error
}
