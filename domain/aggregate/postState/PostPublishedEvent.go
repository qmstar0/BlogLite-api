package postState

type PostPublishedEvent struct {
	Uid int
	Pid int
}

func (c PostPublishedEvent) Topic() string {
	return ""
}
