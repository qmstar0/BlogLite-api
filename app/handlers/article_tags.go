package handlers

import (
	"blog/app/dto"
	"blog/app/response"
	"blog/app/service"
	"blog/infra/e"
	"blog/router"
	"github.com/gin-gonic/gin"
	"strconv"
)

type tagDTO interface {
	dto.TagsR
	dto.TagsW
}

// Tags 标签
type Tags struct {
	Srv tagDTO
}

// NewTags 文章标签相关
func NewTags() router.Control {
	return &Tags{Srv: service.GetSrv()}
}

func (t Tags) Index(c *gin.Context) {
	var apiC = response.Api{C: c}
	allTag, err := t.Srv.AllTag(c)
	if err != nil {
		apiC.Response(err)
		return
	}
	resp := make(map[string]any)
	resp["items"] = allTag
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
	if err := t.Srv.NewTag(c, tagStore); err != nil {
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
	tag, err := t.Srv.GetTag(c, []int{tIntId})
	if err != nil {
		apiC.Response(err)
		return
	}
	apiC.Success(tag)
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
	if err := t.Srv.UpdateTag(c, tIntId, tagStore); err != nil {
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
	_, err = t.Srv.GetTag(c, []int{tIntId})
	if err != nil {
		apiC.Response(err)
		return
	}
	if err := t.Srv.DeleteTag(c, tIntId); err != nil {
		apiC.Response(err)
		return
	}
	apiC.Success(nil)
}
