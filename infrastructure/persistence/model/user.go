package model

import (
	"gorm.io/gorm"
	"time"
)

type UserM struct {
	Id  int
	Uid int

	Name     string
	Email    string
	Password string

	Role int

	CreateAt int64
	UpdateAt int64
	DeleteAt int64
}

func (*UserM) TableName() string {
	return "blog_user"
}
func (m *UserM) BeforeUpdate(db *gorm.DB) error {
	var err error
	t := time.Now().Unix()
	if err != nil {
		return err
	}
	db.Statement.SetColumn("update_at", t)
	return nil
}

func (m *UserM) BeforeCreate(db *gorm.DB) error {
	var err error
	t := time.Now().Unix()
	if err != nil {
		return err
	}
	db.Statement.SetColumn("create_at", t)
	return nil
}
