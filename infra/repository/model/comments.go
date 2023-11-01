package model

import (
	"blog/domain/articles"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type Comments struct {
	*articles.Comments
}

// MarshalBinary 实现encoding.BinaryMarshaler接口 用于redis存储时序列化存数据
func (c *Comments) MarshalBinary() (data []byte, err error) {
	return json.Marshal(c)
}

// UnmarshalBinary 实现encoding.BinaryUnmarshaler接口 用于redis存储时反序列化取数据
func (c *Comments) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, c)
}

// TableName 实现gorm的Tabler表名接口 数据库表名
func (c *Comments) TableName() string {
	return "blog_comments"
}

// BeforeCreate 实现gorm钩子接口 钩子函数，在创建前
func (c *Comments) BeforeCreate(tx *gorm.DB) error {
	var err error
	t := uint(time.Now().Unix())
	if err != nil {
		return err
	}
	tx.Statement.SetColumn("publish_at", t)
	//tx.Statement.SetColumn("aid", uuid.New().String())
	return nil
}
