package postsTags

type PostTagsUpdated struct {
	Uid       int
	Pid       int
	OldTagIds []int
	NewTagIds []int
}
