package handlers

import (
	"blog/app/dto"
	"blog/app/response"
	"blog/app/service"
	"blog/infra/config"
	"blog/infra/e"
	"blog/router"
	"blog/utils"
	"github.com/gin-gonic/gin"
)

type articleDTO interface {
	dto.ArticleR
	dto.TagsR
	dto.CateR
	dto.ArticleW
	dto.TagsW
	dto.CateW
	dto.UserR
}

// Article 文章
type Article struct {
	Srv articleDTO
}

// NewArticle 文章处理相关
func NewArticle() router.Control {
	return &Article{Srv: service.GetSrv()}
}

// NewDraft 草稿相关
func NewDraft() router.Draft {
	return &Article{Srv: service.GetSrv()}
}

// NewTrash 回收站相关
func NewTrash() router.Trash {
	return &Article{Srv: service.GetSrv()}
}

// NewImgUpload 图片上传相关
func NewImgUpload() router.Img {
	return &Article{Srv: service.GetSrv()}
}

func (a Article) Index(c *gin.Context) {
	var (
		apiC = response.Api{C: c}
		resp = new(dto.ArticleListViews)
	)
	queryLimit := c.DefaultQuery("n", config.Conf.Article.DefaultLimit)
	queryPage := c.DefaultQuery("p", "1")
	limit, offset := utils.Offset(queryPage, queryLimit)
	artList, err := a.Srv.GetArticleDetailList(c, limit, offset, false, false)
	if err != nil {
		if e.Compare(err, e.ItemNotExist) {
			apiC.Success(resp)
			return
		}
		apiC.Response(err)
		return
	}
	for _, art := range artList {
		user, err := a.Srv.GetUserByUid(c, art.Article.Uid)
		if err != nil {
			continue
		}
		art.Author = user
	}
	resp.Items = artList
	apiC.Success(resp)
}

func (a Article) Create(c *gin.Context) {
	var (
		apiC = response.Api{C: c}
		resp = new(dto.CateTagIndexViews)
	)
	allTag, err := a.Srv.AllTag(c)
	if err != nil {
		apiC.Response(err)
		return
	}
	allCate, err := a.Srv.AllCate(c)
	if err != nil {
		apiC.Response(err)
		return
	}
	resp.Tag, resp.Cate, resp.ImgUploadUrl = allTag, allCate, "ImgUploadURL"
	apiC.Success(resp)
}

func (a Article) Store(c *gin.Context) {
	var (
		apiC = response.Api{C: c}
	)
	userId := c.GetString("userId")
	if userId == "" {
		apiC.Response(e.NewError(e.InvalidParam, nil))
		return
	}
	user, err := a.Srv.GetUserByUid(c, userId)
	if err != nil {
		apiC.Response(err)
		return
	}
	store, exists := c.Get("store")
	if !exists {
		apiC.Response(e.NewError(e.InvalidParam, nil))
		return
	}
	articleStore, ok := store.(dto.ArticleStore)
	if !ok {
		apiC.Response(e.NewError(e.InvalidParam, nil))
		return
	}
	err = a.Srv.NewArticle(c, user.Uid, articleStore)
	if err != nil {
		apiC.Response(err)
		return
	}
	apiC.Success(nil)
}

func (a Article) Edit(c *gin.Context) {
	var (
		apiC = response.Api{C: c}
		resp = new(dto.ArticleDetailViews)
	)
	artId := c.Param("aid")
	userId := c.GetString("userId")
	if userId == "" {
		apiC.Response(e.NewError(e.InvalidParam, nil))
		return
	}
	user, err := a.Srv.GetUser(c, userId)
	if err != nil {
		apiC.Response(err)
		return
	}
	artDetail, err := a.Srv.GetArticleDetail(c, artId)
	if err != nil {
		apiC.Response(err)
		return
	}
	allTag, err := a.Srv.AllTag(c)
	if err != nil {
		apiC.Response(err)
		return
	}
	allCate, err := a.Srv.AllCate(c)
	if err != nil {
		apiC.Response(err)
		return
	}
	resp.Tags, resp.Cate, resp.ImgUploadUrl = allTag, allCate, "ImgUploadURL"
	artDetail.Author = user
	resp.Article = artDetail
	apiC.Success(resp)
}

func (a Article) Update(c *gin.Context) {
	var (
		apiC = response.Api{C: c}
	)
	artId := c.Param("aid")
	store, exists := c.Get("store")
	if !exists {
		apiC.Response(e.NewError(e.InvalidParam, nil))
		return
	}
	articleStore, ok := store.(dto.ArticleStore)
	if !ok {
		apiC.Response(e.NewError(e.InvalidParam, nil))
		return
	}
	err := a.Srv.UpdateArticle(c, artId, articleStore)
	if err != nil {
		apiC.Response(err)
		return
	}
	apiC.Success(nil)
}

func (a Article) Destroy(c *gin.Context) {
	var apiC = response.Api{C: c}
	artId := c.Param("aid")
	err := a.Srv.DeleteArtcle(c, artId)
	if err != nil {
		apiC.Response(err)
		return
	}
	apiC.Success(nil)
}

func (a Article) TrashIndex(c *gin.Context) {
	var (
		apiC = response.Api{C: c}
		resp = new(dto.ArticleListViews)
	)
	queryLimit := c.DefaultQuery("n", config.Conf.Article.DefaultLimit)
	queryPage := c.DefaultQuery("p", "1")
	limit, offset := utils.Offset(queryPage, queryLimit)
	artTrashList, err := a.Srv.GetArticleDetailList(c, limit, offset, false, true)
	if err != nil {
		apiC.Response(err)
		return
	}
	resp.Items = artTrashList
	apiC.Success(resp)
}

func (a Article) UnTrash(c *gin.Context) {
	var apiC = response.Api{C: c}
	artId := c.Param("aid")
	if err := a.Srv.PublishArticle(c, artId); err != nil {
		apiC.Response(err)
		return
	}
	apiC.Success(nil)
}

func (a Article) DraftIndex(c *gin.Context) {
	var (
		apiC = response.Api{C: c}
		resp = new(dto.ArticleListViews)
	)
	queryLimit := c.DefaultQuery("n", config.Conf.Article.DefaultLimit)
	queryPage := c.DefaultQuery("p", "1")
	limit, offset := utils.Offset(queryPage, queryLimit)
	artDraftList, err := a.Srv.GetArticleDetailList(c, limit, offset, true, false)
	if err != nil {
		apiC.Response(err)
		return
	}
	resp.Items = artDraftList
	apiC.Success(resp)
}

func (a Article) Publish(c *gin.Context) {
	var apiC = response.Api{C: c}
	artId := c.Param("aid")
	if err := a.Srv.PublishArticle(c, artId); err != nil {
		apiC.Response(err)
		return
	}
	apiC.Success(nil)
}

func (a Article) ImgUpload(c *gin.Context) {
	//TODO implement me
	return
}
