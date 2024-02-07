package category

import (
	"time"
)

type Category struct {
	Cid         int
	Name        Name
	DisplayName string
	SeoDesc     string
	Num         uint
	DeleteAt    int64
}

func CreateCategory(name Name, displayName, seoDesc string) *Category {
	return &Category{
		Cid:         0,
		Name:        name,
		DisplayName: displayName,
		SeoDesc:     seoDesc,
		Num:         0,
		DeleteAt:    0,
	}
}
func (c *Category) ChangeCategory(name Name, displayName, seoDesc string) {
	c.Name = name
	c.SeoDesc = seoDesc
	c.DisplayName = displayName
}

func (c *Category) Delete() error {
	c.DeleteAt = time.Now().Unix()
	return nil
}
