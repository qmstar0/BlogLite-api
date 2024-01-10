package repositoryImpl

import (
	"blog/domain/aggregate/userRole"
	"blog/infrastructure/persistence/mapper"
	"blog/infrastructure/persistence/model"
	"gorm.io/gorm"
)

type UserRoleRepositoryImpl struct {
	DB *gorm.DB
}

func (u UserRoleRepositoryImpl) Save(role userRole.UserRole) error {
	roleModel := mapper.UserRoleDomainToModel(role)
	return u.DB.Save(roleModel).Error
}

func (u UserRoleRepositoryImpl) FindByUid(uid int) (userRole.UserRole, error) {
	var user = &model.UserM{}
	result := u.DB.Model(&model.UserM{}).Where("uid = ?", uid).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return mapper.UserRoleModelToDomain(user), nil
}

func (u UserRoleRepositoryImpl) FindByEmail(email string) (userRole.UserRole, error) {
	var user = &model.UserM{}
	result := u.DB.Model(&model.UserM{}).Where("email = ?", email).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return mapper.UserRoleModelToDomain(user), nil
}
