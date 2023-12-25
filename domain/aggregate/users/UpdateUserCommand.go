package users

type UpdateUserCommand struct {
	Uid  string
	Name string
}

func (u UpdateUserCommand) Topic() string {
	//TODO implement me
	panic("implement me")
}
