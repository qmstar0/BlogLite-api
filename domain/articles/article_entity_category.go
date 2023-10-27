package articles

const (
	defaultCategory = "默认分类"
)

// ArticleCategory 文章分类
type ArticleCategory struct {
	Id          int    `json:"-" gorm:"primaryKey"`
	Name        string `json:"name" gorm:"column:name; type:varchar(10); uniqueIndex; not null"`
	DisplayName string `json:"display_name" gorm:"column:display_name; type:varchar(10)"`
	SeoDesc     string `json:"seo_desc" gorm:"column:seo_desc"`
	ParentId    int    `json:"parent_id" gorm:"column:parent_id"`
	CreateAt    uint   `json:"create_at" gorm:"column:create_at"`
	UpdateAt    uint   `json:"update_at" gorm:"column:update_at; default:0"`
}

func NewArticleCategory(name string, displayName string, seoDesc string) *ArticleCategory {
	return &ArticleCategory{Name: name, DisplayName: displayName, SeoDesc: seoDesc}
}

func (ac *ArticleCategory) SetParentId(id int) {
	ac.ParentId = id
}
