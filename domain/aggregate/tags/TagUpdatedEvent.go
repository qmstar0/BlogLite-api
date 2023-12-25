package tags

type TagUpdatedEvent struct {
	Uid         int
	TagId       int
	Name        string
	DisplayName string
	SeoDesc     string
}

func (c TagUpdatedEvent) Topic() string {
	return ""
}
