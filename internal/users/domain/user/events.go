package user

const (
	Login uint16 = iota
	Logout
	RegistrationSuccess
	PasswordReset
	RolesChanged
	NameChanged
)

type UsernameChanged struct {
	Uid     uint32
	OldName string
	NewName string
}

type UserRegistrationSuccess struct {
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
type UserRolesChanged struct {
	Uid       uint32
	OldRights uint16
	NewRights uint16
}

type UserLogin struct {
	Uid   uint32
	Email string
}

type UserLogout struct {
	Uid   uint32
	Email string
}
