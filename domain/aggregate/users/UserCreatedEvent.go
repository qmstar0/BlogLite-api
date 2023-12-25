package users

type UserCreatedEvent struct {
	Uid   int
	Email string
}

func (c UserCreatedEvent) Topic() string {
	return ""
}
