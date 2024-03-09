package command

import (
	"common/e"
	"common/handler"
	"context"
	"users/domain/user"
)

type ResetPassword struct {
	Uid      uint32
	Passowrd string
}

type ResetPasswordHandler handler.CommandHandler[ResetPassword]

type resetPasswordHandler struct {
	userRepo user.UserRepository
}

func NewResetPasswordHandler(repository user.UserRepository) ResetPasswordHandler {
	return &resetPasswordHandler{userRepo: repository}
}

func (h resetPasswordHandler) Handle(ctx context.Context, cmd ResetPassword) error {
	password, err := user.NewPassword(cmd.Passowrd)
	if err != nil {
		return e.Wrap(e.NewValueObjectErr, err)
	}

	u, err := h.userRepo.Find(ctx, cmd.Uid)
	if err != nil {
		return e.Wrap(e.FindErr, err)
	}

	u.ResetPassowrd(password)

	return h.userRepo.Save(ctx, u)
}
