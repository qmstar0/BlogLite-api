package postsCateogry

type PostCategoryUpdatedEvent struct {
	Uid           int
	Pid           int
	OldCategoryId int
	NewCategoryID int
}
