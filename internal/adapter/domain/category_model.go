package domain

import (
	"github.com/uptrace/bun"
	"go-blog-ddd/internal/application/query"
	"go-blog-ddd/internal/domain/aggregates"
	"go-blog-ddd/internal/domain/values"
)

type CategoryM struct {
	bun.BaseModel `bun:"table:category,alias:category"`
	ID            uint32 `bun:"id,pk"`
	Name          string `bun:"name"`
	Desc          string `bun:"desc"`
	Num           uint32 `bun:"num,scanonly"`
}

func CategoryAggregateToDBModel(category *aggregates.Category) *CategoryM {
	return &CategoryM{
		ID:   category.ID,
		Name: category.Name.String(),
		Desc: category.Desc,
	}
}

func CategoryModelToAggregate(m *CategoryM) *aggregates.Category {
	return &aggregates.Category{
		ID:   m.ID,
		Name: values.CategoryName(m.Name),
		Desc: m.Desc,
	}
}

func CategoryModelToView(m *CategoryM) query.CategoryView {
	return query.CategoryView{
		ID:   m.ID,
		Name: m.Name,
		Desc: m.Desc,
		Num:  m.Num,
	}
}
