package user

type UserName string

func NewUserName(s string) (UserName, error) {
	return UserName(s), nil
}

func (n UserName) String() string {
	return string(n)
}
