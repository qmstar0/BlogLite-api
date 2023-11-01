package model

import (
	"blog/domain/articles"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type ArticleCate struct {
	*articles.ArticleCategory
}

// MarshalBinary 实现encoding.BinaryMarshaler接口 用于redis存储时序列化存数据
func (ac *ArticleCate) MarshalBinary() (data []byte, err error) {
	return json.Marshal(ac)
}

// UnmarshalBinary 实现encoding.BinaryUnmarshaler接口 用于redis存储时反序列化取数据
func (ac *ArticleCate) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, ac)
}

// TableName 实现gorm的Tabler表名接口 数据库表名
func (ac *ArticleCate) TableName() string {
	return "blog_article_category"
}

// BeforeCreate 实现gorm钩子接口 钩子函数，在创建前
func (ac *ArticleCate) BeforeCreate(tx *gorm.DB) error {
	var err error
	t := uint(time.Now().Unix())
	if err != nil {
		return err
	}
	tx.Statement.SetColumn("create_at", t)
	return nil
}

// BeforeUpdate 实现gorm钩子接口 钩子函数，在更新前
func (ac *ArticleCate) BeforeUpdate(tx *gorm.DB) error {
	var err error
	t := uint(time.Now().Unix())
	if err != nil {
		return err
	}
	tx.Statement.SetColumn("update_at", t)
	return nil
}
