package command

import (
	"common/e"
	"common/handler"
	"context"
	"users/domain/user"
)

type UpdateUsername struct {
	Uid  uint32
	Name string
}

type UpdateUsernameHandler handler.CommandHandler[UpdateUsername]

type updateUsernameHandler struct {
	userRepo user.UserRepository
}

func NewUpdateUsernameHandler(repository user.UserRepository) UpdateUsernameHandler {
	return &updateUsernameHandler{userRepo: repository}
}

func (h updateUsernameHandler) Handle(ctx context.Context, cmd UpdateUsername) error {
	name, err := user.NewUserName(cmd.Name)
	if err != nil {
		return e.Wrap(e.NewValueObjectErr, err)
	}

	u, err := h.userRepo.Find(ctx, cmd.Uid)
	if err != nil {
		return e.Wrap(e.FindErr, err)
	}

	u.ChangeUsername(name)

	return h.userRepo.Save(ctx, u)
}
