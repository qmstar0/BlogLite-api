package entity

type Category interface {
	Update(name, displayName, seoDesc string) error
	SetParent(id uint) error
	Delete() error
}

type CategoryImpl struct {
	Id          int    `json:"-" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"column:name; type:varchar(10); uniqueIndex; not null"`
	DisplayName string `json:"display_name" gorm:"column:display_name; type:varchar(10)"`
	SeoDesc     string `json:"seo_desc" gorm:"column:seo_desc"`
	ParentId    int    `json:"parent_id" gorm:"column:parent_id"`
	CreateAt    uint   `json:"create_at" gorm:"column:create_at"`
	UpdateAt    uint   `json:"update_at" gorm:"column:update_at; default:0"`
}

func (c *CategoryImpl) Update(name, displayName, seoDesc string) error {
	c.Name = name
	c.DisplayName = displayName
	c.SeoDesc = seoDesc
	return nil
}

func (c *CategoryImpl) SetParent(id uint) error {
	//TODO implement me
	panic("implement me")
}

func (c *CategoryImpl) Delete() error {
	//TODO implement me
	panic("implement me")
}
