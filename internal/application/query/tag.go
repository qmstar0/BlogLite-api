package query

import (
	"context"
)

type TagView struct {
	Name string `json:"name"`
	Num  int    `json:"num"`
}

type TagListView struct {
	Count int       `json:"count"`
	Items []TagView `json:"items"`
}

type TagQueryControl interface {
	All(ctx context.Context) (TagListView, error)
}
