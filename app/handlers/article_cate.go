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

type cateDTO interface {
	dto.CateW
	dto.CateR
}

// Category 分类
type Category struct {
	Srv cateDTO
}

// NewCate 文章分类相关
func NewCate() router.Control {
	return &Category{Srv: service.GetSrv()}
}

func (ac Category) Index(c *gin.Context) {
	var apiC = response.Api{C: c}
	allCate, err := ac.Srv.AllCate(c)
	if err != nil {
		apiC.UnavailableFailResp(err)
		return
	}
	resp := make(map[string]any)
	resp["items"] = allCate
	apiC.Success(resp)
}

func (ac Category) Create(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (ac Category) Store(c *gin.Context) {
	var apiC = response.Api{C: c}
	store, exists := c.Get("store")
	if !exists {
		apiC.ValidateFailResp(e.NewError(e.InvalidParam, nil))
		return
	}
	cateStore, ok := store.(dto.CateStore)
	if !ok {
		apiC.ValidateFailResp(e.NewError(e.InvalidParam, nil))
		return
	}
	if err := ac.Srv.NewCate(c, cateStore); err != nil {
		apiC.UnavailableFailResp(err)
		return
	}
	apiC.Success(nil)
}

func (ac Category) Edit(c *gin.Context) {
	var apiC = response.Api{C: c}
	cid := c.Param("cid")
	cIntId, err := strconv.Atoi(cid)
	if err != nil {
		apiC.ValidateFailResp(err)
		return
	}
	cate, err := ac.Srv.GetCate(c, cIntId)
	if err != nil {
		apiC.NotFoundFailResp(err)
		return
	}
	apiC.Success(cate)
}

func (ac Category) Update(c *gin.Context) {
	var apiC = response.Api{C: c}
	cid := c.Param("cid")
	cIntId, err := strconv.Atoi(cid)
	if err != nil {
		apiC.ValidateFailResp(err)
		return
	}
	store, exists := c.Get("store")
	if !exists {
		apiC.ValidateFailResp(e.NewError(e.InvalidParam, nil))
		return
	}
	cateStore, ok := store.(dto.CateStore)
	if !ok {
		apiC.ValidateFailResp(e.NewError(e.InvalidParam, nil))
		return
	}
	if err := ac.Srv.UpdateCate(c, cIntId, cateStore); err != nil {
		apiC.UnavailableFailResp(err)
		return
	}
	apiC.Success(nil)
}

func (ac Category) Destroy(c *gin.Context) {
	var apiC = response.Api{C: c}
	cid := c.Param("cid")
	cIntId, err := strconv.Atoi(cid)
	if err != nil {
		apiC.ValidateFailResp(err)
		return
	}
	_, err = ac.Srv.GetCate(c, cIntId)
	if err != nil {
		apiC.NotFoundFailResp(err)
		return
	}
	_, err = ac.Srv.GetCateByParentId(c, cIntId)
	if !e.Compare(err, e.ItemNotExist) {
		apiC.ConflictFailResp(e.NewError(e.InvalidDelete, err))
		return
	}
	if err := ac.Srv.DeleteCate(c, cIntId); err != nil {
		apiC.UnavailableFailResp(err)
		return
	}
	apiC.Success(nil)
}
