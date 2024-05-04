package commands

type CreatePost struct {
	Uri        string
	MDFilePath string
}
type DeletePost struct {
	ID uint32
}
type ResetPostCategory struct {
	ID         uint32
	CategoryID uint32
}

type ModifyPostVisibility struct {
	ID      uint32
	Visible bool
}

type ModifyPostTags struct {
	ID      uint32
	NewTags []string
}

type ModifyPostInfo struct {
	ID    uint32
	Title string
	Desc  string
}

type ResetPostContent struct {
	ID         uint32
	MDFilePath string
}
