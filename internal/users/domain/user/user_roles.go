package user

const Normally UserRole = 0

const (
	NewPost UserRole = 1 << iota
)

type UserRole uint16

func IsAdmin(roles UserRole) bool {
	if roles != 0 && roles&(roles<<1+1) == roles {
		return true
	}
	return false
}
