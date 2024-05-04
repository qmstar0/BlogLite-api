package values

type PostUri string

func NewPostUri(s string) (PostUri, error) {
	return PostUri(s), nil
}

func (p PostUri) String() string {
	return string(p)
}
