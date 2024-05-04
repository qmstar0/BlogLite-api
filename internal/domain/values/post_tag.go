package values

type Tag string

func NewTag(s string) (Tag, error) {
	return Tag(s), nil
}

func (t Tag) String() string {
	return string(t)
}
