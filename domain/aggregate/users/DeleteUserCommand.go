package users

type DeleteUserCommand struct {
	Uid   int
	Email string
}

func (d DeleteUserCommand) Topic() string {
	//TODO implement me
	panic("implement me")
}
