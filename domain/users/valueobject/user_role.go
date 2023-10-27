package valueobject

import (
	"database/sql/driver"
)

const (
	CommonUserTag = 1 // 普通用户
	SubscriberTag = 2 // 订阅者
	PublisherTag  = 4 // 发布者

	AdminTag = 15 // 管理员
)

type Role uint

// NewUserRole 新建用户角色
func NewUserRole(i uint) Role {
	return Role(i)
}

// Scan Scanner接口 用于将数据库数据写入结构体
func (ur *Role) Scan(src any) error {
	*ur = Role(src.(int64))
	return nil
}

// Value Valuer接口 用于结构体数据写入数据库
func (ur *Role) Value() (driver.Value, error) {
	return int64(*ur), nil
}

// IsCommonUser 是普通用户
func (ur *Role) IsCommonUser() bool {
	return CommonUserTag&(*ur) == CommonUserTag
}

// IsSubscriber 是订阅者
func (ur *Role) IsSubscriber() bool {
	return SubscriberTag&(*ur) == SubscriberTag
}

// IsPublisher 是发布者
func (ur *Role) IsPublisher() bool {
	return PublisherTag&(*ur) == PublisherTag
}

// IsAdmin 是管理员
func (ur *Role) IsAdmin() bool {
	return AdminTag&(*ur) == AdminTag
}

func (ur *Role) ToUint() uint {
	return uint(*ur)
}

// Add 添加用户标识
func (ur *Role) Add(tag uint) Role {
	u := uint(*ur)
	u |= tag
	*ur = Role(u)
	return *ur
}

// Remove 删除用户标识
func (ur *Role) Remove(tag uint) Role {
	u := uint(*ur)
	u &^= tag
	*ur = Role(u)
	return *ur
}
