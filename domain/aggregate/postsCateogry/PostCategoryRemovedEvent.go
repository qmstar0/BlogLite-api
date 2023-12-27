package postsCateogry

type PostCategoryRemovedEvent struct {
	Uid           int
	Pid           int
	OldCategoryId int
}
