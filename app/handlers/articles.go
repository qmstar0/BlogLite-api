package handlers

import (
	"blog/app/dto"
	"blog/app/response"
	"blog/domain/articles"
	"blog/domain/articles/valueobject"
	"blog/domain/users"
	"blog/infra/config"
	"blog/infra/e"
	"blog/router"
	"blog/utils"
	"github.com/gin-gonic/gin"
)

// Article 文章
type Article struct {
}

// NewArticle 文章处理相关
func NewArticle() router.Control {
	return &Article{}
}

// NewStatus 文章状态
func NewStatus() router.Statue {
	return &Article{}
}

// NewTrash 回收站相关

// NewImgUpload 图片上传相关
func NewImgUpload() router.Img {
	return &Article{}
}

func (a Article) Index(c *gin.Context) {
	var (
		err  error
		apiC = response.Api{C: c}
		resp = new(dto.ArticleListViews)
	)
	queryLimit := c.DefaultQuery("n", config.Conf.Article.DefaultLimit)
	queryPage := c.DefaultQuery("p", "1")
	limit, offset := utils.Offset(queryPage, queryLimit)
	Articles, err := Dao.AllArticle(c, limit, offset, 0b0010)
	if err != nil {
		if e.Compare(err, e.ItemNotExist) {
			apiC.Success(resp)
			return
		}
		apiC.Response(err)
		return
	}
	var (
		artDetails = make([]dto.ArticleListDisplay, len(Articles))
		artD       dto.ArticleDisplay
		cateD      dto.CateDisplay
	)
	for i, art := range Articles {
		artD = dto.ArticleDisplay{
			Aid:       art.Aid,
			Uid:       art.Uid,
			TitleSlug: art.TiTleSlug.ToString(),
			Title:     art.Title,
			Summary:   art.Summary,
			Content:   art.Content,
			PublishAt: art.PublishAt,
			CreateAt:  art.CreateAt,
			UpdateAt:  art.UpdateAt,
			DeleteAt:  art.DeleteAt,
		}
		tags, err := Dao.GetTag(c, art.TagIds)
		if err != nil {
			apiC.Response(err)
			return
		}
		var tagD = make([]dto.TagDisplay, len(tags))
		for j, tag := range tags {
			tagD[j] = dto.TagDisplay{
				Id:          0,
				Name:        tag.Name,
				DisplayName: tag.DisplayName,
				SeoDesc:     tag.SeoDesc,
				Num:         tag.Num,
			}
		}

		cate, err := Dao.GetCate(c, &articles.ArticleCategory{Id: art.CategoryId})
		if err != nil {
			apiC.Response(err)
			return
		}
		cateD = dto.CateDisplay{
			Id:          cate.Id,
			Name:        cate.Name,
			DisplayName: cate.DisplayName,
			SeoDesc:     cate.SeoDesc,
		}
		artDetails[i] = dto.ArticleListDisplay{
			Article:  artD,
			Tags:     tagD,
			Category: cateD,
			Author:   dto.UserDisplay{},
		}
	}
	resp.Items = artDetails

	apiC.Success(resp)
}

func (a Article) Create(c *gin.Context) {
	var (
		apiC = response.Api{C: c}
		resp = new(dto.CateTagIndexViews)
	)
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
	resp.Tag, resp.Cate, resp.ImgUploadUrl = allTagD, allCateD, "ImgUploadURL"
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
	user, err := Dao.GetUser(c, &users.User{Uid: userId})
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
	articleMate := articles.NewArticleMate(user.Uid)
	err = articleMate.SetTitleSlug(articleStore.TitleSlug)
	if err != nil {
		apiC.Response(err)
		return
	}
	articleMate.SetTitle(articleStore.Title)
	articleMate.SetSummary(articleStore.Summary)
	articleMate.SetStatus(valueobject.Draft)
	articleMate.SetCategory(articleStore.Category)
	articleMate.SetTagIDs(articleStore.Tags)
	if err = Dao.NewArticle(c, articleMate); err != nil {
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
	user, err := Dao.GetUser(c, &users.User{Uid: userId})
	if err != nil {
		apiC.Response(err)
		return
	}
	artDetail, err := Dao.GetArticle(c, &articles.ArticleMate{Aid: artId})
	if err != nil {
		apiC.Response(err)
		return
	}
	tags, err := Dao.GetTag(c, artDetail.TagIds)
	if err != nil {
		apiC.Response(err)
		return
	}
	var TagD = make([]dto.TagDisplay, len(tags))
	for i, tag := range tags {
		TagD[i] = dto.TagDisplay{
			Id:          tag.Id,
			Name:        tag.Name,
			DisplayName: tag.DisplayName,
			SeoDesc:     tag.SeoDesc,
			Num:         tag.Num,
		}
	}
	cate, err := Dao.GetCate(c, &articles.ArticleCategory{Id: artDetail.CategoryId})
	if err != nil {
		apiC.Response(err)
		return
	}
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
	resp.Tags, resp.Cate, resp.ImgUploadUrl = allTagD, allCateD, "ImgUploadURL"
	resp.Article = dto.ArticleListDisplay{
		Article: dto.ArticleDisplay{
			Aid:       artDetail.Aid,
			Uid:       artDetail.Uid,
			TitleSlug: artDetail.TiTleSlug.ToString(),
			Title:     artDetail.Title,
			Summary:   artDetail.Summary,
			Content:   artDetail.Content,
			PublishAt: artDetail.PublishAt,
			CreateAt:  artDetail.CreateAt,
			UpdateAt:  artDetail.UpdateAt,
			DeleteAt:  artDetail.DeleteAt,
			Views:     artDetail.Views,
		},
		Tags: TagD,
		Category: dto.CateDisplay{
			Id:          cate.Id,
			Name:        cate.Name,
			DisplayName: cate.DisplayName,
			SeoDesc:     cate.SeoDesc,
		},
		Author: dto.UserDisplay{
			Uid:   user.Uid,
			Name:  user.UserName,
			Email: user.Email.ToString(),
			Role:  user.Role.ToUint(),
		},
	}
	apiC.Success(resp)
}

func (a Article) Update(c *gin.Context) {
	var (
		err  error
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
	articleMate := &articles.ArticleMate{Aid: artId}
	err = articleMate.SetTitleSlug(articleStore.TitleSlug)
	if err != nil {
		apiC.Response(err)
		return
	}
	articleMate.SetTitle(articleStore.Title)
	articleMate.SetSummary(articleStore.Summary)
	articleMate.SetStatus(valueobject.Draft)
	articleMate.SetCategory(articleStore.Category)
	articleMate.SetTagIDs(articleStore.Tags)
	if err = Dao.UptArticle(c, articleMate); err != nil {
		apiC.Response(err)
		return
	}
	apiC.Success(nil)
}

func (a Article) Destroy(c *gin.Context) {
	var apiC = response.Api{C: c}
	artId := c.Param("aid")
	err := Dao.DelArticle(c, &articles.ArticleMate{Aid: artId})
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
	Articles, err := Dao.AllArticle(c, limit, offset, 0b0100)
	if err != nil {
		if e.Compare(err, e.ItemNotExist) {
			apiC.Success(resp)
			return
		}
		apiC.Response(err)
		return
	}
	var (
		artDetails = make([]dto.ArticleListDisplay, len(Articles))
		artD       dto.ArticleDisplay
		cateD      dto.CateDisplay
	)
	for i, art := range Articles {
		artD = dto.ArticleDisplay{
			Aid:       art.Aid,
			Uid:       art.Uid,
			TitleSlug: art.TiTleSlug.ToString(),
			Title:     art.Title,
			Summary:   art.Summary,
			Content:   art.Content,
			PublishAt: art.PublishAt,
			CreateAt:  art.CreateAt,
			UpdateAt:  art.UpdateAt,
			DeleteAt:  art.DeleteAt,
		}
		tags, err := Dao.GetTag(c, art.TagIds)
		if err != nil {
			apiC.Response(err)
			return
		}
		var tagD = make([]dto.TagDisplay, len(tags))
		for j, tag := range tags {
			tagD[j] = dto.TagDisplay{
				Id:          0,
				Name:        tag.Name,
				DisplayName: tag.DisplayName,
				SeoDesc:     tag.SeoDesc,
				Num:         tag.Num,
			}
		}

		cate, err := Dao.GetCate(c, &articles.ArticleCategory{Id: art.CategoryId})
		if err != nil {
			apiC.Response(err)
			return
		}
		cateD = dto.CateDisplay{
			Id:          cate.Id,
			Name:        cate.Name,
			DisplayName: cate.DisplayName,
			SeoDesc:     cate.SeoDesc,
		}
		artDetails[i] = dto.ArticleListDisplay{
			Article:  artD,
			Tags:     tagD,
			Category: cateD,
			Author:   dto.UserDisplay{},
		}
	}
	resp.Items = artDetails

	apiC.Success(resp)
}

func (a Article) DraftIndex(c *gin.Context) {
	var (
		apiC = response.Api{C: c}
		resp = new(dto.ArticleListViews)
	)
	queryLimit := c.DefaultQuery("n", config.Conf.Article.DefaultLimit)
	queryPage := c.DefaultQuery("p", "1")
	limit, offset := utils.Offset(queryPage, queryLimit)
	Articles, err := Dao.AllArticle(c, limit, offset, 0b0001)
	if err != nil {
		if e.Compare(err, e.ItemNotExist) {
			apiC.Success(resp)
			return
		}
		apiC.Response(err)
		return
	}
	var (
		artDetails = make([]dto.ArticleListDisplay, len(Articles))
		artD       dto.ArticleDisplay
		cateD      dto.CateDisplay
	)
	for i, art := range Articles {
		artD = dto.ArticleDisplay{
			Aid:       art.Aid,
			Uid:       art.Uid,
			TitleSlug: art.TiTleSlug.ToString(),
			Title:     art.Title,
			Summary:   art.Summary,
			Content:   art.Content,
			PublishAt: art.PublishAt,
			CreateAt:  art.CreateAt,
			UpdateAt:  art.UpdateAt,
			DeleteAt:  art.DeleteAt,
		}
		tags, err := Dao.GetTag(c, art.TagIds)
		if err != nil {
			apiC.Response(err)
			return
		}
		var tagD = make([]dto.TagDisplay, len(tags))
		for j, tag := range tags {
			tagD[j] = dto.TagDisplay{
				Id:          0,
				Name:        tag.Name,
				DisplayName: tag.DisplayName,
				SeoDesc:     tag.SeoDesc,
				Num:         tag.Num,
			}
		}

		cate, err := Dao.GetCate(c, &articles.ArticleCategory{Id: art.CategoryId})
		if err != nil {
			apiC.Response(err)
			return
		}
		cateD = dto.CateDisplay{
			Id:          cate.Id,
			Name:        cate.Name,
			DisplayName: cate.DisplayName,
			SeoDesc:     cate.SeoDesc,
		}
		artDetails[i] = dto.ArticleListDisplay{
			Article:  artD,
			Tags:     tagD,
			Category: cateD,
			Author:   dto.UserDisplay{},
		}
	}
	resp.Items = artDetails

	apiC.Success(resp)
}

func (a Article) Publish(c *gin.Context) {
	var apiC = response.Api{C: c}
	artId := c.Param("aid")
	art := &articles.ArticleMate{Aid: artId}
	art.Status.SetPublished()
	if err := Dao.UptArticle(c, art); err != nil {
		apiC.Response(err)
		return
	}
	apiC.Success(nil)
}

func (a Article) ImgUpload(c *gin.Context) {
	//TODO implement me
	return
}
