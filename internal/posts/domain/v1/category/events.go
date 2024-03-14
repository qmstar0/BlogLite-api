package category

const (
	AutoSave uint16 = iota
	Created
	SeoDescChanged
	Deleted
)

type CategoryCretaed struct {
	Name    string
	SeoDesc string
}

type CategorySeoDescChanged struct {
	ID  uint32
	Old string
	New string
}

type CategoryDeleted struct {
	ID uint32
}
