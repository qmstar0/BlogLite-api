package postCateogry

type PostCategoryRemovedEvent struct {
	Uid           int
	Pid           int
	OldCategoryId int
}
