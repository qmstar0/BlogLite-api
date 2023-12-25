package tags

type TagCreatedEvent struct {
	Uid         int
	Name        string
	DisplayName string
	SeoDesc     string
}

func (c TagCreatedEvent) Topic() string {
	return ""
}
