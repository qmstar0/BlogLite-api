package post

type PostRepository interface {
	New(post Post) error
	Save(post Post) error
	FindByPid(pid string) (Post, error)
	FindBySlug(Slug string) (Post, error)
}
