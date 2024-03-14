package command

import (
	"common/e"
	"common/handler"
	"context"
	"users/domain/user"
)

type ResetPassword struct {
	Token    string
	Passowrd string
}

type ResetPasswordHandler handler.CommandHandler[ResetPassword]

type resetPasswordHandler struct {
	userRepo    user.UserRepository
	userAuthSer user.UserAuthService
}

func NewResetPasswordHandler(repository user.UserRepository, service user.UserAuthService) ResetPasswordHandler {
	return &resetPasswordHandler{userRepo: repository, userAuthSer: service}
}

func (h resetPasswordHandler) Handle(ctx context.Context, cmd ResetPassword) error {

	tokenU, err := h.userAuthSer.ParseToken(cmd.Token)
	if err != nil {
		return e.Wrap(e.AuthenticationErr, err)
	}

	password, err := user.NewPassword(cmd.Passowrd)
	if err != nil {
		return e.Wrap(e.NewValueObjectErr, err)
	}

	u, err := h.userRepo.Find(ctx, tokenU.Uid)
	if err != nil {
		return e.Wrap(e.FindEventErr, err)
	}

	u.ResetPassowrd(password)

	return h.userRepo.Save(ctx, u)
}
