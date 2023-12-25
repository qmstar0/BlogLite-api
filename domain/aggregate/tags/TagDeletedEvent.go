package tags

type TagDeletedEvent struct {
	Uid   int
	TagId int
}

func (c TagDeletedEvent) Topic() string {
	return ""
}
