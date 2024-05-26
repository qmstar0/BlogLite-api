package aggregates

import (
	"context"
	"github.com/qmstar0/domain/internal/domain/values"
)

type Category struct {
	AggregateRoot

	ID   uint32
	Name values.CategoryName
	Desc string
}

func NewCategory(cid uint32, name values.CategoryName, desc string) *Category {
	return &Category{
		ID:   cid,
		Name: name,
		Desc: desc,
	}
}

func (c *Category) SeoSeoDesc(desc string) {
	c.Desc = desc
}

func (c *Category) Delete() {}

type CategoryRepository interface {
	NextID(ctx context.Context) (uint32, error)
	FindByID(ctx context.Context, id uint32) (*Category, error)
	FindByName(ctx context.Context, name values.CategoryName) (*Category, error)
	FindOrErrByID(ctx context.Context, id uint32) (*Category, error)
	Save(ctx context.Context, category *Category) error
	Delete(ctx context.Context, cid uint32) error
	ResourceUniquenessCheck(ctx context.Context, name values.CategoryName) error
}
