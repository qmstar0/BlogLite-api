package command

import (
	"common/e"
	"common/handler"
	"context"
	"users/domain/user"
)

type CreateUser struct {
	Name     string
	Email    string
	Password string
}

type CreateUserHandler handler.CommandHandler[CreateUser]

type createUserHandler struct {
	userRepo user.UserRepository
}

func NewCreateUserHandler(repository user.UserRepository) CreateUserHandler {
	return &createUserHandler{userRepo: repository}
}

func (c createUserHandler) Handle(ctx context.Context, cmd CreateUser) error {

	email, err := user.NewEmail(cmd.Email)
	if err != nil {
		return e.Wrap(e.NewValueObjectErr, err)
	}

	uid := email.ToID()

	if exist, err := c.userRepo.Exist(ctx, uid); err != nil {
		return e.Wrap(e.FindEventErr, err)
	} else if exist {
		return e.Wrap(e.ResourceCreated, e.ResourceAlreadyExists)
	}

	name, err := user.NewUserName(cmd.Name)
	if err != nil {
		return e.Wrap(e.NewValueObjectErr, err)
	}

	pwd, err := user.NewPassword(cmd.Password)
	if err != nil {
		return e.Wrap(e.NewValueObjectErr, err)
	}

	return c.userRepo.Save(ctx, user.NewUser(name, email, pwd))
}
