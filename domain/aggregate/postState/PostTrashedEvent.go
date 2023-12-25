package postState

type PostTrashedEvent struct {
	Uid int
	Pid int
}

func (c PostTrashedEvent) Topic() string {
	return ""
}
