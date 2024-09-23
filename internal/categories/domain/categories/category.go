package categories

import (
	"github.com/qmstar0/BlogLite-api/internal/common/e"
)

type Category struct {
	slug        string
	name        string
	description string
}

func NewCategory(slug, name, description string) *Category {
	return &Category{
		slug:        slug,
		name:        name,
		description: description,
	}
}

func (c *Category) ModifyDescription(newDescription string) error {
	if newDescription == "" {
		return e.InvalidActionError("内容不能为空")
	}
	c.description = newDescription
	return nil
}

func (c *Category) Slug() string {
	return c.slug
}

func (c *Category) Name() string {
	return c.name
}

func (c *Category) Description() string {
	return c.description
}

func UnmarshalCategoryFromDatabase(
	slug, name, description string,
) *Category {
	return &Category{
		slug:        slug,
		name:        name,
		description: description,
	}
}
