package domain

import (
	"github.com/uptrace/bun"
	"go-blog-ddd/internal/apps/query"
)

type PostTagM struct {
	bun.BaseModel `bun:"table:post_tag,alias:post_tag"`
	ID            uint32 `bun:"id,pk,autoincrement"`
	PostID        uint32 `bun:"post_id"`
	Tag           string `bun:"tag"`
	Num           int    `bun:"num,scanonly"`
}

func PostTagsModelToView(m []*PostTagM) query.TagListView {
	postTagLen := len(m)
	result := query.TagListView{
		Count: postTagLen,
		Items: make([]query.TagView, postTagLen),
	}
	if postTagLen <= 0 {
		return result
	}

	for i, tagM := range m {
		result.Items[i] = query.TagView{
			Name: tagM.Tag,
			Num:  tagM.Num,
		}
	}

	return result
}
