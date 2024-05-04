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

type CategoryReadModel interface {
	All(context.Context) ([]CategoryView, error)
}

type CategoryQueryControl struct {
	readmodel CategoryReadModel
}

func NewCategoryQueryControl(model CategoryReadModel) CategoryQueryControl {
	return CategoryQueryControl{
		readmodel: model,
	}
}

func (c CategoryQueryControl) GetCategorys(ctx context.Context) (CategoryListView, error) {
	all, err := c.readmodel.All(ctx)
	if err != nil {
		return CategoryListView{}, err
	}
	return CategoryListView{
		Count: len(all),
		Items: all,
	}, nil
}
