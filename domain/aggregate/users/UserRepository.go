package users

type UserRepository interface {
	Save(user User) error
	FindByUid(uid int) (User, error)
	FindByEmail(email string) (User, error)
}
