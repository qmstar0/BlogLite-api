package postCateogry

type PostCategoryRepository interface {
	Save(postCateogry PostCategory) error
	Find(pid int) (PostCategory, error)
}
