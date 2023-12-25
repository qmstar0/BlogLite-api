package users

type UserResetPassowrdCommand struct {
	Uid         int
	NewPassword string
}

func (u UserResetPassowrdCommand) Topic() string {
	//TODO implement me
	panic("implement me")
}
