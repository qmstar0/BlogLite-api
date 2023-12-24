package postTags

type UpdatePostTagsCommand struct {
	Uid    int
	Pid    int
	TagIds []int
}
