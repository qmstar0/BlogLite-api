package postState

type PostRestoredEvent struct {
	Pid int
	Uid int
}

func (c PostRestoredEvent) Topic() string {
	return ""
}
