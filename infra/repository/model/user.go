package model

import (
	"blog/domain/users"
	"encoding/json"
	"gorm.io/gorm"
	"time"
)

type User struct {
	*users.User
}

// MarshalBinary 实现encoding.BinaryMarshaler接口 用于redis存储时序列化存数据
func (u *User) MarshalBinary() (data []byte, err error) {
	return json.Marshal(u)
}

// UnmarshalBinary 实现encoding.BinaryUnmarshaler接口 用于redis存储时反序列化取数据
func (u *User) UnmarshalBinary(data []byte) error {
	return json.Unmarshal(data, u)
}

// TableName 实现gorm的Tabler表名接口 数据库表名
func (*User) TableName() string {
	return "blog_user"
}

// BeforeCreate 实现gorm钩子接口 钩子函数，在创建前
func (u *User) BeforeCreate(tx *gorm.DB) error {
	var err error
	createTime := uint(time.Now().Unix())
	if err != nil {
		return err
	}
	tx.Statement.SetColumn("create_at", createTime)
	return nil
}

// BeforeUpdate 实现gorm钩子接口 钩子函数，在更新前
func (u *User) BeforeUpdate(tx *gorm.DB) error {
	var err error
	t := uint(time.Now().Unix())
	if err != nil {
		return err
	}
	tx.Statement.SetColumn("update_at", t)
	return nil
}
