package query

import (
	"common/e"
	"common/handler"
	"context"
	"users/domain/user"
)

type GetUserAuthenticationHandler handler.QueryHandler[string, string]

type getUserAuthenticationHandler struct {
	userRepo    user.UserRepository
	userAuthSer user.UserAuthService
}

func NewGetUserAuthenticationHandler(repository user.UserRepository, service user.UserAuthService) GetUserAuthenticationHandler {
	return &getUserAuthenticationHandler{
		userRepo:    repository,
		userAuthSer: service,
	}
}

func (g getUserAuthenticationHandler) Handle(ctx context.Context, emailStr string) (string, error) {
	email, err := user.NewEmail(emailStr)
	if err != nil {
		return "", e.Wrap(e.NewValueObjectErr, err)
	}

	u, err := g.userRepo.Find(ctx, email.ToID())
	if err != nil {
		return "", e.Wrap(e.FindEntityErr, err)
	}

	token, err := g.userAuthSer.SignToken(u)
	if err != nil {
		return "", e.Wrap(e.IssueTokenErr, err)
	}

	return token, nil
}
