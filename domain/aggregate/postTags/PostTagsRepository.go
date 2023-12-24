package postTags

type PostTagsRepository interface {
	Save(postTags PostTags) error
	Find(pid int) (PostTags, error)
}
