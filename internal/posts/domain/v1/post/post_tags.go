package post

import (
	"common/idtools"
	"errors"
)

type TagGroup struct {
	Tags     []uint32
	CheckBit uint32
}

func NewTagGroup(tag []string) (TagGroup, error) {
	tagslen := len(tag)
	if tagslen <= 0 {
		return TagGroup{}, nil
	}
	if tagslen > 4 {
		return TagGroup{}, errors.New("a post can only have up to 4 tag")
	}
	group := TagGroup{
		Tags:     make([]uint32, tagslen),
		CheckBit: 0,
	}
	for i := range tagslen {
		hashID := idtools.NewHashID([]byte(tag[i]))
		group.CheckBit += hashID
		group.Tags[i] = hashID
	}
	return group, nil
}
