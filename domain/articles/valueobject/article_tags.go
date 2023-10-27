package valueobject

import (
	"database/sql/driver"
	"encoding/json"
)

type TagS []int

// Scan Scanner接口 用于将数据库数据写入结构体
func (t *TagS) Scan(src any) error {
	bytes := src.([]byte)
	return json.Unmarshal(bytes, t)
}

// Value Valuer接口 用于结构体数据写入数据库
func (t *TagS) Value() (driver.Value, error) {
	return json.Marshal(t)
}
