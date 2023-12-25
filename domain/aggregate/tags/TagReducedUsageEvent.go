package tags

type TagReducedUsageEvent struct {
	Uid   int
	TagId int
}

func (c TagReducedUsageEvent) Topic() string {
	return ""
}
