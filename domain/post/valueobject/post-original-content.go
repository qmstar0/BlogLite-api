package valueobject

type Original string

func NewOriginal(original string) (Original, error) {
	return Original(original), nil
}
