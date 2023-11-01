package dao

import (
	"blog/domain/users"
	"blog/infra/e"
	"blog/infra/repository/model"
	"context"
)

func (d *Dao) NewUser(c context.Context, user *users.User) error {
	var (
		userModel = &model.User{User: user}
	)
	if err := d.db.WithContext(c).Model(&model.User{}).Create(userModel).Error; err != nil {
		return e.NewError(e.DBCreateErr, err)
	}
	return nil
}

func (d *Dao) UptUser(c context.Context, user *users.User) error {
	var (
		err       error
		userModel = &model.User{User: user}
	)
	err = d.db.WithContext(c).Model(&model.User{}).Where("user_email = ? OR uid = ?", userModel.Email, userModel.Uid).Updates(userModel).Error
	if err != nil {
		return e.NewError(e.DBUpdateErr, err)
	}
	return nil
}

func (d *Dao) DelUser(c context.Context, user *users.User) error {
	var (
		userModel = &model.User{User: user}
	)
	if err := d.db.WithContext(c).Model(&model.User{}).Where(userModel).Delete(userModel).Error; err != nil {
		return e.NewError(e.DBDeleteErr, err)
	}
	return nil
}

func (d *Dao) GetUser(c context.Context, user *users.User) (*users.User, error) {
	var (
		userModel = &model.User{User: user}
	)
	result := d.db.WithContext(c).Model(&model.User{}).Where(userModel).Limit(1).Find(userModel)
	if result.Error != nil {
		return nil, e.NewError(e.DBFindErr, result.Error)
	} else if result.RowsAffected == 0 {
		return nil, e.NewError(e.ItemNotExist, nil)
	}
	return userModel.User, nil
}
