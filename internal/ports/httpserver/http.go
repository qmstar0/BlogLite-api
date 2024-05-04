package httpserver

import (
	"github.com/gin-gonic/gin"
	"go-blog-ddd/config"
	"go-blog-ddd/internal/application"
	"go-blog-ddd/internal/ports/httpserver/render"
	"strings"
)

type HttpServer struct {
	app *application.App
}

func NewHttpServer(app *application.App) *HttpServer {
	return &HttpServer{app: app}
}

func (h HttpServer) GetApiCategorys(c *gin.Context) {
	categorys, err := h.app.Queries.Categorys.GetCategorys(c)
	if err != nil {
		render.Error(c, err)
		return
	}
	render.Respond(c, categorys)
}

func (h HttpServer) GetApiPosts(c *gin.Context, params GetApiPostsParams) {
	var (
		page       = 1
		categroyID uint32
		tags       []string
	)

	if params.Page != nil {
		page = *params.Page
	}

	if params.Tag != nil {
		tags = strings.Split(*params.Tag, ",")
	}

	if params.Category != nil {
		categroyID = *params.Category
	}

	result, err := h.app.Queries.Posts.GetPostsWithFilter(c,
		config.Conf.Request.Post.DefaultLimit, page, tags, categroyID)
	if err != nil {
		render.Error(c, err)
		return
	}
	render.Respond(c, result)
}

func (h HttpServer) GetApiPostsId(c *gin.Context, id uint32) {
	result, err := h.app.Queries.Posts.FindByID(c, id)
	if err != nil {
		render.Error(c, err)
		return
	}
	render.Respond(c, result)
}

func (h HttpServer) GetApiTags(c *gin.Context) {
	result, err := h.app.Queries.Tags.GetTags(c)
	if err != nil {
		render.Error(c, err)
		return
	}
	render.Respond(c, result)
}
