package categorys

type CategoryCreatedEvent struct {
	Uid         int
	Name        string
	DisplayName string
	SeoDesc     string
}

func (c CategoryCreatedEvent) Topic() string {
	return ""
}
