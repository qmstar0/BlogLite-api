package categorys

type CategoryDeletedEvent struct {
	Uid        int
	CategoryId int
}

func (c CategoryDeletedEvent) Topic() string {
	return ""
}
