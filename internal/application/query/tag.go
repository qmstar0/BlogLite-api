package query

import (
	"context"
)

type TagListView struct {
	Count int      `json:"count"`
	Items []string `json:"items"`
}

type TagReadModel interface {
	All(ctx context.Context) ([]string, error)
}

type TagQueryControl struct {
	readmodel TagReadModel
}

func NewTagQueryControl(readmodel TagReadModel) TagQueryControl {
	return TagQueryControl{readmodel: readmodel}
}

func (t TagQueryControl) GetTags(ctx context.Context) (TagListView, error) {
	all, err := t.readmodel.All(ctx)
	if err != nil {
		return TagListView{}, err
	}
	return TagListView{
		Count: len(all),
		Items: all,
	}, nil
}
