package handlers

import (
	"blog/app/dto"
	"blog/app/response"
	"blog/domain/articles"
	"blog/infra/e"
	"blog/router"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Tags 标签
type Tags struct {
}

// NewTags 文章标签相关
func NewTags() router.Control {
	return &Tags{}
}

func (t Tags) Index(c *gin.Context) {
	var apiC = response.Api{C: c}
	allTag, err := Dao.AllTag(c)
	if err != nil {
		apiC.Response(err)
		return
	}
	var allTagD = make([]dto.TagDisplay, len(allTag))
	for i, tag := range allTag {
		allTagD[i] = dto.TagDisplay{
			Id:          tag.Id,
			Name:        tag.Name,
			DisplayName: tag.DisplayName,
			SeoDesc:     tag.SeoDesc,
			Num:         tag.Num,
		}
	}
	resp := make(map[string]any)
	resp["items"] = allTagD
	apiC.Success(resp)
}

func (t Tags) Create(c *gin.Context) {
	panic("implement me")
}

func (t Tags) Store(c *gin.Context) {
	var apiC = response.Api{C: c}
	store, exists := c.Get("store")
	if !exists {
		apiC.Response(e.NewError(e.InvalidParam, nil))
		return
	}
	tagStore, ok := store.(dto.TagStore)
	if !ok {
		apiC.Response(e.NewError(e.InvalidParam, nil))
		return
	}
	tag := articles.NewArticleTags(tagStore.Name, tagStore.DisplayName, tagStore.SeoDesc)
	if err := Dao.NewTag(c, tag); err != nil {
		apiC.Response(err)
		return
	}
	apiC.Success(nil)
}

func (t Tags) Edit(c *gin.Context) {
	var apiC = response.Api{C: c}
	tid := c.Param("tid")
	tIntId, err := strconv.Atoi(tid)
	if err != nil {
		apiC.Response(err)
		return
	}
	tags, err := Dao.GetTag(c, []int{tIntId})
	if err != nil {
		apiC.Response(err)
		return
	}
	var allTagD = make([]dto.TagDisplay, len(tags))
	for i, tag := range tags {
		allTagD[i] = dto.TagDisplay{
			Id:          tag.Id,
			Name:        tag.Name,
			DisplayName: tag.DisplayName,
			SeoDesc:     tag.SeoDesc,
			Num:         tag.Num,
		}
	}
	apiC.Success(allTagD)
}

func (t Tags) Update(c *gin.Context) {
	var apiC = response.Api{C: c}
	tid := c.Param("tid")
	tIntId, err := strconv.Atoi(tid)
	if err != nil {
		apiC.Response(err)
		return
	}
	store, exists := c.Get("store")
	if !exists {
		apiC.Response(e.NewError(e.InvalidParam, nil))
		return
	}
	tagStore, ok := store.(dto.TagStore)
	if !ok {
		apiC.Response(e.NewError(e.InvalidParam, nil))
		return
	}
	if err := Dao.UptTag(c, &articles.ArticleTags{
		Id:          tIntId,
		Name:        tagStore.Name,
		DisplayName: tagStore.DisplayName,
		SeoDesc:     tagStore.SeoDesc,
	}); err != nil {
		apiC.Response(err)
		return
	}
	apiC.Success(nil)
}

func (t Tags) Destroy(c *gin.Context) {
	var apiC = response.Api{C: c}
	tid := c.Param("tid")
	tIntId, err := strconv.Atoi(tid)
	if err != nil {
		apiC.Response(err)
		return
	}
	_, err = Dao.GetTag(c, []int{tIntId})
	if err != nil {
		apiC.Response(err)
		return
	}
	if err := Dao.DelTag(c, &articles.ArticleTags{Id: tIntId}); err != nil {
		apiC.Response(err)
		return
	}
	apiC.Success(nil)
}
