package model

import (
	"gorm.io/gorm"
	"time"
)

type CategoryM struct {
	Id       int    `gorm:"primaryKey"`
	Name     string `gorm:"column:name;"`
	Display  string `gorm:"column:display_name"`
	SeoDesc  string `gorm:"column:seo_desc"`
	Num      uint   `gorm:"column:num"`
	CreateAt int64  `gorm:"column:create_at;"`
	DeleteAt int64  `gorm:"column:delete_at;default:0"`
	UpdateAt int64  `gorm:"update_at;default:0"`
}

func (*CategoryM) TableName() string {
	return "blog_category"
}

func (m *CategoryM) BeforeUpdate(db *gorm.DB) error {
	var err error
	t := time.Now().Unix()
	if err != nil {
		return err
	}
	db.Statement.SetColumn("update_at", t)
	return nil
}

func (m *CategoryM) BeforeCreate(db *gorm.DB) error {
	var err error
	t := time.Now().Unix()
	if err != nil {
		return err
	}
	db.Statement.SetColumn("create_at", t)
	return nil
}
