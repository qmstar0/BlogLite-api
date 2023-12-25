package posts

type PostDeletedEvent struct {
	Uid int
	Pid int
}

func (c PostDeletedEvent) Topic() string {
	return ""
}
