package commands

type CreatePost struct {
	Uri    string
	MDFile []byte
}

type DeletePost struct {
	ID uint32
}

type ModifyPost struct {
	ID         uint32
	Tags       *[]string
	CategoryID *uint32
	Visible    *bool
	Title      *string
	Desc       *string
}

type ResetPostContent struct {
	ID     uint32
	MDFile []byte
}
