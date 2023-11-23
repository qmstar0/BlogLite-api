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

// Category 分类
type Category struct {
}

// NewCate 文章分类相关
func NewCate() router.Control {
	return &Category{}
}

func (ac Category) Index(c *gin.Context) {
	var apiC = response.Api{C: c}
	allCate, err := Dao.AllCate(c)
	if err != nil {
		apiC.Response(err)
		return
	}
	var allCateD = make([]dto.CateDisplay, len(allCate))
	for i, cate := range allCate {
		allCateD[i] = dto.CateDisplay{
			Id:          cate.Id,
			Name:        cate.Name,
			DisplayName: cate.DisplayName,
			SeoDesc:     cate.SeoDesc,
		}
	}
	resp := make(map[string]any)
	resp["items"] = allCateD
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
		apiC.Response(e.NewError(e.InvalidParam, nil))
		return
	}
	cateStore, ok := store.(dto.CateStore)
	if !ok {
		apiC.Response(e.NewError(e.InvalidParam, nil))
		return
	}
	category := articles.NewArticleCategory(cateStore.Name, cateStore.DisplayName, cateStore.SeoDesc)
	if category.ParentId != 0 {
		category.SetParentId(category.ParentId)
	}
	if err := Dao.NewCate(c, category); err != nil {
		apiC.Response(err)
		return
	}
	apiC.Success(nil)
}

func (ac Category) Edit(c *gin.Context) {
	var apiC = response.Api{C: c}
	cid := c.Param("cid")
	cIntId, err := strconv.Atoi(cid)
	if err != nil {
		apiC.Response(err)
		return
	}
	cate, err := Dao.GetCate(c, &articles.ArticleCategory{Id: cIntId})
	if err != nil {
		apiC.Response(err)
		return
	}
	apiC.Success(cate)
}

func (ac Category) Update(c *gin.Context) {
	var (
		err  error
		apiC = response.Api{C: c}
	)
	cid := c.Param("cid")
	cIntId, err := strconv.Atoi(cid)
	if err != nil {
		apiC.Response(err)
		return
	}
	store, exists := c.Get("store")
	if !exists {
		apiC.Response(e.NewError(e.InvalidParam, nil))
		return
	}
	cateStore, ok := store.(dto.CateStore)
	if !ok {
		apiC.Response(e.NewError(e.InvalidParam, nil))
		return
	}
	if err = Dao.UptCate(c, &articles.ArticleCategory{
		Id:          cIntId,
		Name:        cateStore.Name,
		DisplayName: cateStore.DisplayName,
		SeoDesc:     cateStore.SeoDesc,
		ParentId:    cateStore.ParentId,
	}); err != nil {
		apiC.Response(err)
		return
	}
	apiC.Success(nil)
}

func (ac Category) Destroy(c *gin.Context) {
	var apiC = response.Api{C: c}
	cid := c.Param("cid")
	cIntId, err := strconv.Atoi(cid)
	if err != nil {
		apiC.Response(err)
		return
	}
	_, err = Dao.GetCate(c, &articles.ArticleCategory{Id: cIntId})
	if err != nil {
		apiC.Response(err)
		return
	}
	_, err = Dao.GetCate(c, &articles.ArticleCategory{ParentId: cIntId})
	if !e.Compare(err, e.ItemNotExist) {
		apiC.Response(e.NewError(e.InvalidDelete, err))
		return
	}
	if err := Dao.DelCate(c, &articles.ArticleCategory{Id: cIntId}); err != nil {
		apiC.Response(err)
		return
	}
	apiC.Success(nil)
}
