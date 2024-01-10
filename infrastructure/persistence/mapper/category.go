package mapper

import (
	"blog/domain/aggregate/categorys"
	"blog/infrastructure/persistence/model"
)

func CategoryDoaminToModel(category categorys.Category) *model.CategoryM {
	c := category.(*categorys.CategoryImpl)
	return &model.CategoryM{
		Id:       c.Id,
		Name:     c.Name,
		Display:  c.DisplayName,
		SeoDesc:  c.SeoDesc,
		Num:      c.Num,
		DeleteAt: c.DeleteAt,
	}
}

func CategoryModelToDomain(category *model.CategoryM) *categorys.CategoryImpl {
	return &categorys.CategoryImpl{
		Id:          category.Id,
		Name:        category.Name,
		DisplayName: category.Display,
		SeoDesc:     category.SeoDesc,
		Num:         category.Num,
		DeleteAt:    category.DeleteAt,
	}
}
