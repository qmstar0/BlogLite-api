package tags

type TagIncreasedUsageEvent struct {
	Uid   int
	TagId int
}

func (c TagIncreasedUsageEvent) Topic() string {
	return ""
}
