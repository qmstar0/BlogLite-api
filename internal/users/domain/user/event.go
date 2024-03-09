package user

const (
	PasswordReset = iota + 10
	RightsChanged
)

type UsernameUpdated struct {
	Uid     uint32
	OldName string
	NewName string
}

type UserCreated struct {
	Uid      uint32
	Name     string
	Email    string
	Password string
	Rights   uint16
}
type UserPasswordReset struct {
	Uid         uint32
	OldPassowrd string
	NewPassword string
}
type UserRightsChanged struct {
	Uid       uint32
	OldRights uint16
	NewRights uint16
}
