package users

import (
	"blog/domain/common"
	"time"
)

type User interface {
	ResetPassword(pub common.DomainEventPublisher, cmd UserResetPassowrdCommand) error
	Update(pub common.DomainEventPublisher, cmd UpdateUserCommand) error
	Delete(pub common.DomainEventPublisher, cmd DeleteUserCommand) error
}

type UserImpl struct {
	Uid      int
	Name     string
	Email    string
	Password string
	DeleteAt int64
}

func CreateUser(pub common.DomainEventPublisher, cmd CreateUserCommand) (User, error) {
	uid := GenUId.NextID()
	if err := pub.Publish(UserCreatedEvent{
		Uid:   uid,
		Email: cmd.Email,
	}); err != nil {
		return nil, err
	}
	return &UserImpl{
		Uid:   uid,
		Email: cmd.Email,
	}, nil
}
func (u *UserImpl) ResetPassword(pub common.DomainEventPublisher, cmd UserResetPassowrdCommand) error {
	u.Password = cmd.NewPassword
	return pub.Publish(UserPasswordResetedEvent{
		Uid:         u.Uid,
		NewPassword: cmd.NewPassword,
	})
}

func (u *UserImpl) Update(pub common.DomainEventPublisher, cmd UpdateUserCommand) error {
	u.Name = cmd.Name
	return pub.Publish(UserUpdatedEvent{
		Uid:  u.Uid,
		Name: cmd.Name,
	})
}

func (u *UserImpl) Delete(pub common.DomainEventPublisher, cmd DeleteUserCommand) error {
	u.DeleteAt = time.Now().Unix()
	return pub.Publish(UserDeletedEvent{
		Uid:   u.Uid,
		Email: cmd.Email,
	})
}
