package repositoryImpl

import (
	"blog/domain/aggregate/users"
	"blog/infrastructure/persistence/mapper"
	"blog/infrastructure/persistence/model"
	"gorm.io/gorm"
)

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func (u UserRepositoryImpl) Save(user users.User) error {
	userModel := mapper.UserDomainToModel(user)
	return u.DB.Save(userModel).Error
}

func (u UserRepositoryImpl) FindByUid(uid int) (users.User, error) {
	var user = &model.UserM{}
	result := u.DB.Model(&model.UserM{}).Where("uid = ?", uid).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return mapper.UserModelToDomain(user), nil
}

func (u UserRepositoryImpl) FindByEmail(email string) (users.User, error) {
	var user = &model.UserM{}
	result := u.DB.Model(&model.UserM{}).Where("email = ?", email).First(user)
	if result.Error != nil {
		return nil, result.Error
	}
	return mapper.UserModelToDomain(user), nil
}
