package articles

import (
	"blog/domain/articles/valueobject"
	"blog/infra/e"
	"blog/utils"
)

// ArticleMate 文章
type ArticleMate struct {
	Id        int    `json:"-" gorm:"primaryKey"`
	Aid       string `json:"aid" gorm:"column:aid; uniqueIndex; not null"`
	Uid       string `json:"uid" gorm:"column:uid; not null"`
	Title     string `json:"title" gorm:"column:title; type:varchar(255)"`
	TiTleSlug string `json:"title_slug" gorm:"column:title_slug; uniqueIndex; type:varchar(255)"`
	Summary   string `json:"summary" gorm:"column:summary"`
	Content   string `json:"content" gorm:"column:content; type:longtext"`
	Original  string `json:"original_content" gorm:"column:original_content; type:longtext"`
	CreateAt  uint   `json:"create_at" gorm:"column:create_at"`
	UpdateAt  uint   `json:"update_at" gorm:"column:update_at; default:0"`
	PublishAt uint   `json:"publish_at" gorm:"column:publish_at; default:0"`
	DeleteAt  uint   `json:"delete_at" gorm:"column:delete_at; default:0"`

	CategoryId int                `json:"category_id" gorm:"column:category_id; type:TINYINT UNSIGNED"`
	Views      uint               `json:"views" gorm:"column:views; default:0"`
	TagIds     valueobject.TagS   `json:"tag_ids" gorm:"column:tag_ids; type:json"`
	Status     valueobject.Status `json:"status" gorm:"column:status; type:TINYINT UNSIGNED"`
}

func NewArticleMate(uid string) *ArticleMate {
	return &ArticleMate{Uid: uid}
}

func (a *ArticleMate) SetStatus(status int64) {
	switch status {
	case valueobject.Draft:
		a.Status.SetDraft()
	case valueobject.Deleted:
		a.Status.SetDeleted()
	case valueobject.Published:
		a.Status.SetPublished()
	}
}

func (a *ArticleMate) SetCategory(cateId int) {
	a.CategoryId = cateId
}

func (a *ArticleMate) SetTagIDs(tagIDs valueobject.TagS) {
	a.TagIds = tagIDs
}

func (a *ArticleMate) SetSummary(summary string) {
	a.Summary = summary
}

func (a *ArticleMate) SetTitle(title string) {
	a.Title = title
}
func (a *ArticleMate) SetTitleSlug(title string) {
	a.TiTleSlug = title
}
func (a *ArticleMate) SetContent(content string) error {
	a.Original = content
	html, err := utils.MarkdownToHTML(content)
	if err != nil {
		return e.NewError(e.MarkdownTOHTMLErr, err)
	}
	a.Content = html
	return nil
}
