package query

import (
	"common/e"
	"common/handler"
	"context"
	"users/domain/user"
)

type GetUser struct {
	Uid uint32
}

type GetUserResult struct {
	Uid    uint32
	Name   string
	Rights uint16
}

type GetUserHandler handler.QueryHandler[GetUser, GetUserResult]

type getUserHandler struct {
	userRepo user.UserRepository
}

func NewGetUserHandler(repository user.UserRepository) GetUserHandler {
	return &getUserHandler{userRepo: repository}
}

func (g getUserHandler) Handle(ctx context.Context, q GetUser) (GetUserResult, error) {
	find, err := g.userRepo.Find(ctx, q.Uid)
	if err != nil {
		return GetUserResult{}, e.Wrap(e.FindErr, err)
	}
	return GetUserResult{
		Uid:    find.Uid,
		Name:   find.Name.String(),
		Rights: uint16(find.Rights),
	}, nil
}
