package articles

// Comments 评论
type Comments struct {
	Id           uint   `json:"-" gorm:"primaryKey"`
	Cid          string `json:"cid" gorm:"column:cid; uniqueIndex; not null"`
	Uid          string `json:"uid" gorm:"column:uid; not null"`
	Aid          string `json:"aid" gorm:"column:aid; not null"`
	Content      string `json:"content" gorm:"column:content; not null"`
	ParentNodeId string `json:"parentnode_id" gorm:"column:parentnode_id"`

	PublishAt uint `json:"publish_at" gorm:"column:publish_at; not null"`
	DeleteAt  uint `json:"delete_at" gorm:"column:delete_at; default:0"`
}

// GetPublishTime 获取发布时间
func (c *Comments) GetPublishTime() uint {
	return c.PublishAt
}

// GetDeleteTime 获取删除时间
func (c *Comments) GetDeleteTime() uint {
	return c.DeleteAt
}

// GetCid 获取cid
func (c *Comments) GetCid() string {
	return c.Cid
}

// GetUid 获取uid
func (c *Comments) GetUid() string {
	return c.Uid
}

// GetAid 获取aid
func (c *Comments) GetAid() string {
	return c.Aid
}

// GetContent 获取内容
func (c *Comments) GetContent() string {
	return c.Content
}

// GetParentNodeId 获取父节点id
func (c *Comments) GetParentNodeId() string {
	return c.ParentNodeId
}
