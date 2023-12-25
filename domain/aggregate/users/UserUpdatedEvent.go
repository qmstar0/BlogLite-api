package users

type UserUpdatedEvent struct {
	Uid  int
	Name string
}

func (c UserUpdatedEvent) Topic() string {
	return ""
}
