package post

const (
	AutoSave uint16 = iota
	Created
	TitleChanged
	ContentChanged
	SeoDescChanged
	StateChanged
	TagsChanged
	CategoryChanged
	Deleted
)

type PostCreated struct {
	UserID uint32
	Uri    string
}

// PostChanged Title/Content/SeoDesc
type PostChanged struct {
	ID  uint32
	Old string
	New string
}

type PostStateChanged struct {
	ID  uint32
	Old uint16
	New uint16
}

type PostTagsChanged struct {
	ID  uint32
	Old []uint32
	New []uint32
}

type PostCategoryChanged struct {
	ID  uint32
	Old uint32
	New uint32
}

type PostDeleted struct {
	ID uint32
}
