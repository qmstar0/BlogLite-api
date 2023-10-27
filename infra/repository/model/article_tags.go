package model

import (
	"blog/domain/articles"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type ArticleTags struct {
	*articles.ArticleTags
}

// MarshalBinary 实现encoding.BinaryMarshaler接口 用于redis存储时序列化存数据
func (at *ArticleTags) MarshalBinary() (data []byte, err error) {
	return json.Marshal(at)
}

// UnmarshalBinary 实现encoding.BinaryUnmarshaler接口 用于redis存储时反序列化取数据
func (at *ArticleTags) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, at)
}

// TableName 实现gorm的Tabler表名接口 数据库表名
func (at *ArticleTags) TableName() string {
	return "blog_article_tags"
}

// BeforeCreate 实现gorm钩子接口 钩子函数，在创建前
func (at *ArticleTags) BeforeCreate(tx *gorm.DB) error {
	var err error
	t := uint(time.Now().Unix())
	if err != nil {
		return err
	}
	tx.Statement.SetColumn("create_at", t)
	return nil
}

// BeforeUpdate 实现gorm钩子接口 钩子函数，在更新前
func (at *ArticleTags) BeforeUpdate(tx *gorm.DB) error {
	var err error
	t := uint(time.Now().Unix())
	if err != nil {
		return err
	}
	tx.Statement.SetColumn("update_at", t)
	return nil
}
