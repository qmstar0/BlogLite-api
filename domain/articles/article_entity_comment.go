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
