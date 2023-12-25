package categorys

type CategoryReducedUsageEvent struct {
	Uid        int
	CategoryId int
}

func (c CategoryReducedUsageEvent) Topic() string {
	return ""
}
