package user

const Normally UserRights = 0

const (
	NewPost UserRights = 1 << iota
)

type UserRights uint16

func IsAdmin(rights UserRights) bool {
	if rights != 0 && rights&(rights<<1+1) == rights {
		return true
	}
	return false
}
