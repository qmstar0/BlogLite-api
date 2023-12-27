package postsTags

type UpdatePostTagsCommand struct {
	Uid    int
	Pid    int
	TagIds []int
}
