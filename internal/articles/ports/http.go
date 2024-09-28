package ports

import (
	"github.com/gin-gonic/gin"
	"github.com/qmstar0/BlogLite-api/internal/articles/application"
	"github.com/qmstar0/BlogLite-api/internal/articles/application/command"
	"github.com/qmstar0/BlogLite-api/internal/articles/application/query"
	"github.com/qmstar0/BlogLite-api/internal/common/auth"
	"github.com/qmstar0/BlogLite-api/internal/common/server/httpresponse"
	"github.com/qmstar0/BlogLite-api/pkg/utils"
	"io"
	"strings"
)

type HttpServer struct {
	app *application.App
}

func NewHttpServer(app *application.App) *HttpServer {
	return &HttpServer{app: app}
}

func (h HttpServer) GetArticleList(c *gin.Context, params GetArticleListParams) {

	var (
		includeInvisible bool
		tags             []string
	)

	if err := auth.FilterAuthWithUserType(c.Request.Context(), "admin"); err == nil {
		if params.IncludeInvisible != nil {
			includeInvisible = *params.IncludeInvisible
		}
	}

	if params.Tags != nil && *params.Tags != "" {
		tags = strings.Split(*params.Tags, ",")
	}

	view, err := h.app.Query.ArticleList.Handle(c.Request.Context(), query.ArticleList{
		Category:         params.Category,
		Tags:             tags,
		Page:             params.Page,
		Limit:            params.Limit,
		IncludeInvisible: includeInvisible,
	})

	if err != nil {
		httpresponse.Error(c, err)
		return
	}

	httpresponse.Data(c, view)
}

func (h HttpServer) InitializationArticle(c *gin.Context) {
	req, err := utils.BindJSON[InitializationArticleJSONRequestBody](c)
	if err != nil {
		httpresponse.Error(c, err)
		return
	}

	err = h.app.Command.InitializationArticle.Handle(c.Request.Context(), command.InitializationArticle{
		Uri:        req.Uri,
		CategoryID: req.Category,
	})

	httpresponse.ErrorOrOK(c, err)
}

func (h HttpServer) DeleteArticle(c *gin.Context, uri string) {
	err := h.app.Command.DeleteArticle.Handle(c.Request.Context(), command.DeleteArticle{Uri: uri})

	httpresponse.ErrorOrOK(c, err)
}

func (h HttpServer) GetArticleDetail(c *gin.Context, uri string, params GetArticleDetailParams) {
	view, err := h.app.Query.ArticleDetail.Handle(c.Request.Context(), query.ArticleDetail{
		URI:     uri,
		Version: params.Version,
	})
	if err != nil {
		httpresponse.Error(c, err)
		return
	}
	httpresponse.Data(c, view)
}

func (h HttpServer) SetArticleVersion(c *gin.Context, uri string) {
	req, err := utils.BindJSON[SetArticleVersionJSONRequestBody](c)
	if err != nil {
		httpresponse.Error(c, err)
		return
	}

	err = h.app.Command.SetArticleVersion.Handle(c.Request.Context(), command.SetArticleVersion{
		Uri:     uri,
		Version: req.Version,
	})

	httpresponse.ErrorOrOK(c, err)
}

func (h HttpServer) ChangeArticleCategory(c *gin.Context, uri string) {
	req, err := utils.BindJSON[ChangeArticleCategoryJSONRequestBody](c)
	if err != nil {
		httpresponse.Error(c, err)
		return
	}
	err = h.app.Command.ChangeArticleCategory.Handle(c.Request.Context(), command.ChangeArticleCategory{
		Uri:        uri,
		CategoryID: req.Category,
	})

	httpresponse.ErrorOrOK(c, err)
}

func (h HttpServer) ModifyArticleTags(c *gin.Context, uri string) {
	req, err := utils.BindJSON[ModifyArticleTagsJSONRequestBody](c)
	if err != nil {
		httpresponse.Error(c, err)
		return
	}
	err = h.app.Command.ModifyArticleTags.Handle(c.Request.Context(), command.ModifyArticleTags{
		Uri:  uri,
		Tags: req.Tags,
	})

	httpresponse.ErrorOrOK(c, err)
}

func (h HttpServer) GetArticleVersion(c *gin.Context, uri string) {
	if err := auth.FilterAuthWithUserType(c.Request.Context(), "admin"); err != nil {
		httpresponse.Error(c, err)
		return
	}

	view, err := h.app.Query.ArticleVersionList.Handle(c.Request.Context(), query.ArticleVersionList{Uri: uri})
	if err != nil {
		httpresponse.Error(c, err)
		return
	}

	httpresponse.Data(c, view)
}

func (h HttpServer) CreateNewArticleVersion(c *gin.Context, uri string) {
	formFile, err := c.FormFile("content")
	if err != nil {
		httpresponse.Error(c, err)
		return
	}
	file, err := formFile.Open()
	if err != nil {
		httpresponse.Error(c, err)
		return
	}

	contentBytes, err := io.ReadAll(file)
	if err != nil {
		httpresponse.Error(c, err)
		return
	}

	err = h.app.Command.AddNewVersion.Handle(c.Request.Context(), command.AddNewVersion{
		Uri:    uri,
		Source: string(contentBytes),
	})

	httpresponse.ErrorOrOK(c, err)
}

func (h HttpServer) RemoveArticleVersion(c *gin.Context, uri string, version string) {
	err := h.app.Command.RemoveVersion.Handle(c.Request.Context(), command.RemoveVersion{
		Uri:     uri,
		Version: version,
	})

	httpresponse.ErrorOrOK(c, err)
}

func (h HttpServer) ChangeArticleVisibility(c *gin.Context, uri string) {
	req, err := utils.BindJSON[ChangeArticleVisibilityJSONRequestBody](c)
	if err != nil {
		httpresponse.Error(c, err)
		return
	}

	err = h.app.Command.ChangeArticleVisibility.Handle(c.Request.Context(), command.ChangeArticleVisibility{
		Uri:        uri,
		Visibility: req.Visibility,
	})

	httpresponse.ErrorOrOK(c, err)
}

func (h HttpServer) GetAllTags(c *gin.Context) {
	view, err := h.app.Query.TagList.Handle(c)
	if err != nil {
		httpresponse.Error(c, err)
		return
	}
	httpresponse.Data(c, view)
}
