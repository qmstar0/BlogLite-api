package values

type PostTitle string

func NewPostTitle(title string) (PostTitle, error) {
	return PostTitle(title), nil
}

func (p PostTitle) String() string {
	return string(p)
}
