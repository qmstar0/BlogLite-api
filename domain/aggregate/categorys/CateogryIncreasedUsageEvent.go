package categorys

type CategoryIncreasedUsageEvent struct {
	Uid        int
	CategoryId int
}

func (c CategoryIncreasedUsageEvent) Topic() string {
	return ""
}
