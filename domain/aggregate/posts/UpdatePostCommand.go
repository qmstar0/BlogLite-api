package posts

type UpdatePostCommand struct {
	Uid     int
	Title   string
	Slug    string
	Summary string
	Content string
}

func (u UpdatePostCommand) Topic() string {
	//TODO implement me
	panic("implement me")
}
