package userRole

type UserRole interface {
	IsPublisher() bool
	IsAdmin() bool
	IsSubscriber() bool
	IsNormal() bool
}

const (
	roleSubscriber = 1 << iota
	rolePublished
)

type UserRoleImpl struct {
	Uid  int
	Role int
}

func (u *UserRoleImpl) IsNormal() bool {
	return u.Role == 0
}

func (u *UserRoleImpl) IsPublisher() bool {
	return u.Role&rolePublished == rolePublished
}

func (u *UserRoleImpl) IsAdmin() bool {
	return u.Role&(u.Role+1) == 0
}

func (u *UserRoleImpl) IsSubscriber() bool {
	return u.Role&roleSubscriber == roleSubscriber
}
