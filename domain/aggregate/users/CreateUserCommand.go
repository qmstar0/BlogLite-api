package users

type CreateUserCommand struct {
	Email string
}

func (c CreateUserCommand) Topic() string {
	//TODO implement me
	panic("implement me")
}
