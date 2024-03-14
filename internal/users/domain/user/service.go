package user

type UserAuthService interface {
	SignToken(user *User) (string, error)
	ParseToken(token string) (*User, error)
	IsAdmin(user *User) bool
}
