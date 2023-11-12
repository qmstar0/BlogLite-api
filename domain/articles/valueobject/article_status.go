package valueobject

import "database/sql/driver"

const (
	Draft = 1 << iota
	Published
	Deleted
	SinglePage
)

var statuMap = map[Status]string{
	Draft:      "Draft",
	Published:  "Published",
	Deleted:    "Deleted",
	SinglePage: "SinglePage",
}

type Status uint

func NewStatus(status uint) Status {
	return Status(status)
}

// Scan Scanner接口 用于将数据库数据写入结构体
func (s *Status) Scan(src any) error {
	status := src.(int64)
	*s = Status(status)
	return nil
}

// Value Valuer接口 用于结构体数据写入数据库
func (s *Status) Value() (driver.Value, error) {
	return int64(*s), nil
}

// String 获取数据代表的含义
func (s *Status) String() string {
	return statuMap[*s]
}

// IsDraft 草稿
func (s *Status) IsDraft() bool {
	return Draft&(*s) == Draft
}

// IsPublished 已发布
func (s *Status) IsPublished() bool {
	return Published&(*s) == Published
}

// IsDeleted 已删除
func (s *Status) IsDeleted() bool {
	return Deleted&(*s) == Deleted
}

// IsSinglePage 是单页
func (s *Status) IsSinglePage() bool {
	return SinglePage&(*s) == SinglePage
}

// SetDraft 设置草稿
func (s *Status) SetDraft() {
	*s = Draft
}

// SetPublished 设置已发布
func (s *Status) SetPublished() {
	*s = Published
}

// SetDeleted 设置已删除
func (s *Status) SetDeleted() {
	*s = Deleted
}

// SetSinglePage 设置单页
func (s *Status) SetSinglePage() {
	*s = SinglePage
}
