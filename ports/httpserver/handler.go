package httpserver

import (
	"fmt"
	"github.com/charmbracelet/log"
	"github.com/gin-gonic/gin"
	"go-blog-ddd/config"
	"go-blog-ddd/internal/apps"
	"go-blog-ddd/internal/apps/service"
	"go-blog-ddd/internal/domain/commands"
	"go-blog-ddd/internal/pkg/e"
	"go-blog-ddd/pkg/logging"
	"go-blog-ddd/ports/httpserver/render"
	"io"
	"strings"
	"time"
)

type HttpServer struct {
	domain *apps.DomainApp
	auth   *service.AdminAuthenticationService
	logger *log.Logger
}

func NewHttpServer(
	domain *apps.DomainApp,
	auth *service.AdminAuthenticationService,
) *HttpServer {
	return &HttpServer{domain: domain, auth: auth, logger: logging.WithPrefix("http handler")}
}

func (h HttpServer) DeleteCategory(c *gin.Context, id uint32) {
	err := h.domain.Commands.DeleteCategory.Handle(c, commands.DeleteCategory{ID: id})
	render.Respond(c, err, nil)
}

func (h HttpServer) ModifyCategoryDesc(c *gin.Context, id uint32) {
	var req ModifyCategoryDescJSONRequestBody

	if err := c.ShouldBindJSON(&req); err != nil {
		render.Error(c, e.PErrInvalidParam.WithError(err))
		return
	}
	err := h.domain.Commands.ModifyCategoryDesc.Handle(c, commands.ModifyCategoryDesc{
		ID:      id,
		NewDesc: req.Desc,
	})
	render.Respond(c, err, nil)
}

func (h HttpServer) GetAllCategorys(c *gin.Context) {
	all, err := h.domain.Queries.Categorys.All(c)
	render.Respond(c, err, all)
}

func (h HttpServer) CreateCategory(c *gin.Context) {
	var req CreateCategoryJSONRequestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		render.Error(c, e.PErrInvalidParam.WithError(err))
		return
	}
	ResourceID, err := h.domain.Commands.CreateCategory.Handle(c, commands.CreateCategory{
		Name: req.Name,
		Desc: req.Desc,
	})
	c.Header("ResourceID", fmt.Sprintf("%d", ResourceID))
	render.Respond(c, err, nil)
}

func (h HttpServer) DeletePost(c *gin.Context, id uint32) {
	err := h.domain.Commands.DeletePost.Handle(c, commands.DeletePost{ID: id})
	render.Respond(c, err, nil)
}

func (h HttpServer) ModifyPosts(c *gin.Context, id uint32) {
	var req ModifyPostsJSONRequestBody
	if err := c.ShouldBindJSON(&req); err != nil {
		render.Error(c, e.PErrInvalidParam.WithError(err))
		return
	}

	if req.Visible == nil && req.Desc == nil && req.Category == nil && req.Tags == nil && req.Title == nil {
		render.Success(c)
		return
	}
	err := h.domain.Commands.ModifyPost.Handle(c, commands.ModifyPost{
		ID:         id,
		Tags:       req.Tags,
		CategoryID: req.Category,
		Visible:    req.Visible,
		Title:      req.Title,
		Desc:       req.Desc,
	})
	render.Respond(c, err, nil)
}

func (h HttpServer) GetPostByUri(c *gin.Context, uri string) {
	result, err := h.domain.Queries.Posts.FindByUri(c, uri)
	render.Respond(c, err, result)
}

func (h HttpServer) GetPostList(c *gin.Context, params GetPostListParams) {
	var (
		page     = 1
		limit    = config.Cfg.HttpRequest.Post.DefaultLimit
		tags     []string
		categroy uint32
	)

	if params.Page != nil {
		if *params.Page > 0 {
			page = *params.Page
		}
	}
	if params.Limit != nil {
		if *params.Limit < 100 && *params.Limit > 0 {
			limit = *params.Limit
		}
	}
	if params.Category != nil {
		categroy = *params.Category
	}

	if params.Tag != nil {
		tags = strings.Split(*params.Tag, ",")
	}

	result, err := h.domain.Queries.Posts.AllWithFilter(
		c,
		limit,
		page,
		tags,
		categroy,
		true,
	)
	render.Respond(c, err, result)
}

func (h HttpServer) CreatePost(c *gin.Context) {
	formfile, err := c.FormFile("content")
	if err != nil {
		render.Error(c, err)
		return
	}

	file, err := formfile.Open()
	if err != nil {
		render.Error(c, err)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		render.Error(c, err)
		return
	}
	ResourceID, err := h.domain.Commands.CreatePost.Handle(c, commands.CreatePost{
		Uri:    c.PostForm("uri"),
		MDFile: content,
	})
	c.Header("ResourceID", fmt.Sprintf("%d", ResourceID))
	render.Respond(c, err, nil)
}

func (h HttpServer) GetRecentPostList(c *gin.Context) {
	result, err := h.domain.Queries.Posts.RecentPosts(c, config.Cfg.HttpRequest.Post.RecentPostsNum)
	render.Respond(c, err, result)
}

func (h HttpServer) GetAllTags(c *gin.Context) {
	result, err := h.domain.Queries.Tags.All(c)
	render.Respond(c, err, result)
}

func (h HttpServer) SetPostContent(c *gin.Context, id uint32) {
	formfile, err := c.FormFile("content")
	if err != nil {
		render.Error(c, err)
		return
	}

	file, err := formfile.Open()
	if err != nil {
		render.Error(c, err)
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		render.Error(c, err)
		return
	}
	err = h.domain.Commands.ResetPostContent.Handle(c, commands.ResetPostContent{
		ID:     id,
		MDFile: content,
	})
	render.Respond(c, err, nil)
}

func (h HttpServer) AdminGetPostList(c *gin.Context, params AdminGetPostListParams) {
	var page = 1

	if params.Page != nil {
		if *params.Page > 0 {
			page = *params.Page
		}
	}

	result, err := h.domain.Queries.Posts.AllWithFilter(
		c,
		config.Cfg.HttpRequest.Post.DefaultLimit,
		page,
		nil,
		0,
		false,
	)
	render.Respond(c, err, result)
}

func (h HttpServer) GetAdminToken(c *gin.Context, params GetAdminTokenParams) {
	if params.Key == nil {
		render.Error(c, e.PErrInvalidParam)
		return
	}
	if *params.Key != config.Cfg.JWTAuth.AuthKey {
		render.Error(c, e.PErrInvalidParam)
		return
	}
	result, err := h.auth.GenerateAdminToken(time.Hour * 24 * time.Duration(config.Cfg.JWTAuth.AuthTokenLife))
	render.Respond(c, err, result)
}
