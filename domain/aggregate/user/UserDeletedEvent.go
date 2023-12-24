package user

type UserDeletedEvent struct {
	Uid   int
	Email string
}
