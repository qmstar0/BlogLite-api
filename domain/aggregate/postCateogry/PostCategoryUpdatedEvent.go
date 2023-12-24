package postCateogry

type PostCategoryUpdatedEvent struct {
	Uid           int
	Pid           int
	OldCategoryId int
	NewCategoryID int
}
