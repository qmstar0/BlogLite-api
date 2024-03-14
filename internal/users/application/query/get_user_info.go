package query

import (
	"common/e"
	"common/handler"
	"context"
	"users/domain/user"
)

type GetUserInfo struct {
	Email string
}

type GetUserInfoResult struct {
	Uid    uint32
	Name   string
	Rights uint16
}

type GetUserInfoHandler handler.QueryHandler[GetUserInfo, GetUserInfoResult]

type getUserInfoHandler struct {
	userRepo user.UserRepository
}

func NewGetUserInfoHandler(repository user.UserRepository) GetUserInfoHandler {
	return &getUserInfoHandler{userRepo: repository}
}

func (g getUserInfoHandler) Handle(ctx context.Context, q GetUserInfo) (GetUserInfoResult, error) {
	email, err := user.NewEmail(q.Email)
	if err != nil {
		return GetUserInfoResult{}, e.Wrap(e.NewValueObjectErr, err)
	}

	find, err := g.userRepo.Find(ctx, email.ToID())
	if err != nil {
		return GetUserInfoResult{}, e.Wrap(e.FindEventErr, err)
	}
	return GetUserInfoResult{
		Uid:    find.Uid,
		Name:   find.Name.String(),
		Rights: uint16(find.Roles),
	}, nil
}
