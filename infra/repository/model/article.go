package model

import (
	"blog/domain/articles"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type ArticleMate struct {
	*articles.ArticleMate
}

// MarshalBinary 实现encoding.BinaryMarshaler接口 用于redis存储时序列化存数据
func (a *ArticleMate) MarshalBinary() (data []byte, err error) {
	return json.Marshal(a)
}

// UnmarshalBinary 实现encoding.BinaryUnmarshaler接口 用于redis存储时反序列化取数据
func (a *ArticleMate) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, a)
}

// TableName 实现gorm的Tabler表名接口 数据库表名
func (*ArticleMate) TableName() string {
	return "blog_articles"
}

// BeforeCreate 实现gorm钩子接口 钩子函数，在创建前
func (a *ArticleMate) BeforeCreate(tx *gorm.DB) error {
	var err error
	t := uint(time.Now().Unix())
	if err != nil {
		return err
	}
	tx.Statement.SetColumn("create_at", t)
	return nil
}

// BeforeUpdate 实现gorm钩子接口 钩子函数，在更新前
func (a *ArticleMate) BeforeUpdate(tx *gorm.DB) error {
	var err error
	t := uint(time.Now().Unix())
	if err != nil {
		return err
	}
	tx.Statement.SetColumn("update_at", t)
	return nil
}
