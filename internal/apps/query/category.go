package query

import (
	"context"
)

type CategoryView struct {
	ID   uint32 `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Desc string `json:"desc,omitempty"`
	Num  uint32 `json:"num,omitempty"`
}

type CategoryListView struct {
	Count int            `json:"count"`
	Items []CategoryView `json:"items"`
}

type CategoryQueryControl interface {
	All(context.Context) (CategoryListView, error)
}
