package articles

type ArticleTags struct {
	Id          int    `json:"-" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"column:name; type:varchar(10); not null"`
	DisplayName string `json:"display_name" gorm:"column:display_name; type:varchar(10)"`
	SeoDesc     string `json:"seo_desc" gorm:"column:seo_desc"`
	Num         uint   `json:"num" gorm:"column:num; not null; default:0"`
	CreateAt    uint   `json:"create_at" gorm:"column:create_at"`
	UpdateAt    uint   `json:"update_at" gorm:"column:update_at; default:0"`
}

func NewArticleTags(name string, displayName string, seoDesc string) *ArticleTags {
	return &ArticleTags{Name: name, DisplayName: displayName, SeoDesc: seoDesc, Num: 0}
}
