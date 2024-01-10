package repositoryImpl

import (
	"blog/domain/aggregate/categorys"
	"blog/infrastructure/persistence/mapper"
	"blog/infrastructure/persistence/model"
	"gorm.io/gorm"
)

type CategoryRepositoryImpl struct {
	DB *gorm.DB
}

func (c CategoryRepositoryImpl) Save(cate categorys.Category) error {
	toModel := mapper.CategoryDoaminToModel(cate)
	return c.DB.Save(toModel).Error
}

func (c CategoryRepositoryImpl) FindById(id int) (categorys.Category, error) {
	var category = &model.CategoryM{}
	result := c.DB.Model(&model.CategoryM{}).Where("id = ?", id).First(category)
	if result.Error != nil {
		return nil, result.Error
	}
	return mapper.CategoryModelToDomain(category), nil
}

func (c CategoryRepositoryImpl) FindByName(name string) (categorys.Category, error) {
	var category = &model.CategoryM{}
	result := c.DB.Model(&model.CategoryM{}).Where("name = ?", name).First(category)
	if result.Error != nil {
		return nil, result.Error
	}
	return mapper.CategoryModelToDomain(category), nil
}
