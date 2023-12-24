package user

type UserPasswordResetedEvent struct {
	Uid         int
	NewPassword string
}
