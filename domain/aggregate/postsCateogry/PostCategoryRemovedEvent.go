package postsCateogry

type PostCategoryRemovedEvent struct {
	Uid           int
	Pid           int
	OldCategoryId int
}

func (c PostCategoryRemovedEvent) Topic() string {
	return ""
}
