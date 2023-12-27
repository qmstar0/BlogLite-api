package users

type UserDeletedEvent struct {
	Uid   int
	Email string
}
