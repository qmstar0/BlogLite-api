package users

type UserPasswordResetedEvent struct {
	Uid         int
	NewPassword string
}
