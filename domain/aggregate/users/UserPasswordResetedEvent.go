package users

type UserPasswordResetedEvent struct {
	Uid         int
	NewPassword string
}

func (c UserPasswordResetedEvent) Topic() string {
	return ""
}
