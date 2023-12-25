package users

type UserDeletedEvent struct {
	Uid   int
	Email string
}

func (c UserDeletedEvent) Topic() string {
	return ""
}
