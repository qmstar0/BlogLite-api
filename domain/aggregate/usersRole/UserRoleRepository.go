package usersRole

type UserRoleRepository interface {
	Save(role UserRole) error
	FindByUid(uid int) (UserRole, error)
	FindByEmail(email string) (UserRole, error)
}
