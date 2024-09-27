// Package ports provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package ports

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
	openapi_types "github.com/oapi-codegen/runtime/types"
)

const (
	BearerScopes = "bearer.Scopes"
)

// GetArticleListParams defines parameters for GetArticleList.
type GetArticleListParams struct {
	Page             *int    `form:"page,omitempty" json:"page,omitempty"`
	Limit            *int    `form:"limit,omitempty" json:"limit,omitempty"`
	Filter           *string `form:"filter,omitempty" json:"filter,omitempty"`
	IncludeInvisible *bool   `form:"includeInvisible,omitempty" json:"includeInvisible,omitempty"`
}

// InitializationArticleJSONBody defines parameters for InitializationArticle.
type InitializationArticleJSONBody struct {
	// Category 分组
	Category string `json:"category"`
	Uri      string `json:"uri"`
}

// GetArticleDetailParams defines parameters for GetArticleDetail.
type GetArticleDetailParams struct {
	Version *string `form:"version,omitempty" json:"version,omitempty"`
}

// SetArticleVersionJSONBody defines parameters for SetArticleVersion.
type SetArticleVersionJSONBody struct {
	Version string `json:"version"`
}

// ChangeArticleCategoryJSONBody defines parameters for ChangeArticleCategory.
type ChangeArticleCategoryJSONBody struct {
	Category string `json:"category"`
}

// ModifyArticleTagsJSONBody defines parameters for ModifyArticleTags.
type ModifyArticleTagsJSONBody struct {
	Tags []string `json:"tags"`
}

// CreateNewArticleVersionMultipartBody defines parameters for CreateNewArticleVersion.
type CreateNewArticleVersionMultipartBody struct {
	Content *openapi_types.File `json:"content,omitempty"`
}

// ChangeArticleVisibilityJSONBody defines parameters for ChangeArticleVisibility.
type ChangeArticleVisibilityJSONBody struct {
	Visibility bool `json:"visibility"`
}

// InitializationArticleJSONRequestBody defines body for InitializationArticle for application/json ContentType.
type InitializationArticleJSONRequestBody InitializationArticleJSONBody

// SetArticleVersionJSONRequestBody defines body for SetArticleVersion for application/json ContentType.
type SetArticleVersionJSONRequestBody SetArticleVersionJSONBody

// ChangeArticleCategoryJSONRequestBody defines body for ChangeArticleCategory for application/json ContentType.
type ChangeArticleCategoryJSONRequestBody ChangeArticleCategoryJSONBody

// ModifyArticleTagsJSONRequestBody defines body for ModifyArticleTags for application/json ContentType.
type ModifyArticleTagsJSONRequestBody ModifyArticleTagsJSONBody

// CreateNewArticleVersionMultipartRequestBody defines body for CreateNewArticleVersion for multipart/form-data ContentType.
type CreateNewArticleVersionMultipartRequestBody CreateNewArticleVersionMultipartBody

// ChangeArticleVisibilityJSONRequestBody defines body for ChangeArticleVisibility for application/json ContentType.
type ChangeArticleVisibilityJSONRequestBody ChangeArticleVisibilityJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// 获取文章列表
	// (GET /articles)
	GetArticleList(c *gin.Context, params GetArticleListParams)
	// 初始化文章
	// (POST /articles)
	InitializationArticle(c *gin.Context)
	// 删除文章
	// (DELETE /articles/{uri})
	DeleteArticle(c *gin.Context, uri string)
	// 获取文章详情
	// (GET /articles/{uri})
	GetArticleDetail(c *gin.Context, uri string, params GetArticleDetailParams)
	// 设置文章当前版本
	// (PUT /articles/{uri}/)
	SetArticleVersion(c *gin.Context, uri string)
	// 更改文章分类
	// (PATCH /articles/{uri}/category)
	ChangeArticleCategory(c *gin.Context, uri string)
	// 修改文章标签
	// (PATCH /articles/{uri}/tags)
	ModifyArticleTags(c *gin.Context, uri string)
	// 获取文章全部版本
	// (GET /articles/{uri}/versions)
	GetArticleVersion(c *gin.Context, uri string)
	// 创建新版本
	// (POST /articles/{uri}/versions)
	CreateNewArticleVersion(c *gin.Context, uri string)
	// 移除文章某版本
	// (DELETE /articles/{uri}/versions/{version})
	RemoveArticleVersion(c *gin.Context, uri string, version string)
	// 修改文章可见性
	// (PATCH /articles/{uri}/visibility)
	ChangeArticleVisibility(c *gin.Context, uri string)
	// 获取所有标签
	// (GET /tags)
	GetAllTags(c *gin.Context)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetArticleList operation middleware
func (siw *ServerInterfaceWrapper) GetArticleList(c *gin.Context) {

	var err error

	c.Set(BearerScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetArticleListParams

	// ------------- Optional query parameter "page" -------------

	err = runtime.BindQueryParameter("form", true, false, "page", c.Request.URL.Query(), &params.Page)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter page: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", c.Request.URL.Query(), &params.Limit)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter limit: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "filter" -------------

	err = runtime.BindQueryParameter("form", true, false, "filter", c.Request.URL.Query(), &params.Filter)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter filter: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Optional query parameter "includeInvisible" -------------

	err = runtime.BindQueryParameter("form", true, false, "includeInvisible", c.Request.URL.Query(), &params.IncludeInvisible)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter includeInvisible: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetArticleList(c, params)
}

// InitializationArticle operation middleware
func (siw *ServerInterfaceWrapper) InitializationArticle(c *gin.Context) {

	c.Set(BearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.InitializationArticle(c)
}

// DeleteArticle operation middleware
func (siw *ServerInterfaceWrapper) DeleteArticle(c *gin.Context) {

	var err error

	// ------------- Path parameter "uri" -------------
	var uri string

	err = runtime.BindStyledParameterWithOptions("simple", "uri", c.Param("uri"), &uri, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter uri: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(BearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteArticle(c, uri)
}

// GetArticleDetail operation middleware
func (siw *ServerInterfaceWrapper) GetArticleDetail(c *gin.Context) {

	var err error

	// ------------- Path parameter "uri" -------------
	var uri string

	err = runtime.BindStyledParameterWithOptions("simple", "uri", c.Param("uri"), &uri, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter uri: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(BearerScopes, []string{})

	// Parameter object where we will unmarshal all parameters from the context
	var params GetArticleDetailParams

	// ------------- Optional query parameter "version" -------------

	err = runtime.BindQueryParameter("form", true, false, "version", c.Request.URL.Query(), &params.Version)
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter version: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetArticleDetail(c, uri, params)
}

// SetArticleVersion operation middleware
func (siw *ServerInterfaceWrapper) SetArticleVersion(c *gin.Context) {

	var err error

	// ------------- Path parameter "uri" -------------
	var uri string

	err = runtime.BindStyledParameterWithOptions("simple", "uri", c.Param("uri"), &uri, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter uri: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(BearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.SetArticleVersion(c, uri)
}

// ChangeArticleCategory operation middleware
func (siw *ServerInterfaceWrapper) ChangeArticleCategory(c *gin.Context) {

	var err error

	// ------------- Path parameter "uri" -------------
	var uri string

	err = runtime.BindStyledParameterWithOptions("simple", "uri", c.Param("uri"), &uri, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter uri: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(BearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ChangeArticleCategory(c, uri)
}

// ModifyArticleTags operation middleware
func (siw *ServerInterfaceWrapper) ModifyArticleTags(c *gin.Context) {

	var err error

	// ------------- Path parameter "uri" -------------
	var uri string

	err = runtime.BindStyledParameterWithOptions("simple", "uri", c.Param("uri"), &uri, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter uri: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(BearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ModifyArticleTags(c, uri)
}

// GetArticleVersion operation middleware
func (siw *ServerInterfaceWrapper) GetArticleVersion(c *gin.Context) {

	var err error

	// ------------- Path parameter "uri" -------------
	var uri string

	err = runtime.BindStyledParameterWithOptions("simple", "uri", c.Param("uri"), &uri, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter uri: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(BearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetArticleVersion(c, uri)
}

// CreateNewArticleVersion operation middleware
func (siw *ServerInterfaceWrapper) CreateNewArticleVersion(c *gin.Context) {

	var err error

	// ------------- Path parameter "uri" -------------
	var uri string

	err = runtime.BindStyledParameterWithOptions("simple", "uri", c.Param("uri"), &uri, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter uri: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(BearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateNewArticleVersion(c, uri)
}

// RemoveArticleVersion operation middleware
func (siw *ServerInterfaceWrapper) RemoveArticleVersion(c *gin.Context) {

	var err error

	// ------------- Path parameter "uri" -------------
	var uri string

	err = runtime.BindStyledParameterWithOptions("simple", "uri", c.Param("uri"), &uri, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter uri: %w", err), http.StatusBadRequest)
		return
	}

	// ------------- Path parameter "version" -------------
	var version string

	err = runtime.BindStyledParameterWithOptions("simple", "version", c.Param("version"), &version, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter version: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(BearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.RemoveArticleVersion(c, uri, version)
}

// ChangeArticleVisibility operation middleware
func (siw *ServerInterfaceWrapper) ChangeArticleVisibility(c *gin.Context) {

	var err error

	// ------------- Path parameter "uri" -------------
	var uri string

	err = runtime.BindStyledParameterWithOptions("simple", "uri", c.Param("uri"), &uri, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter uri: %w", err), http.StatusBadRequest)
		return
	}

	c.Set(BearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ChangeArticleVisibility(c, uri)
}

// GetAllTags operation middleware
func (siw *ServerInterfaceWrapper) GetAllTags(c *gin.Context) {

	c.Set(BearerScopes, []string{})

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetAllTags(c)
}

// GinServerOptions provides options for the Gin server.
type GinServerOptions struct {
	BaseURL      string
	Middlewares  []MiddlewareFunc
	ErrorHandler func(*gin.Context, error, int)
}

// RegisterHandlers creates http.Handler with routing matching OpenAPI spec.
func RegisterHandlers(router gin.IRouter, si ServerInterface) {
	RegisterHandlersWithOptions(router, si, GinServerOptions{})
}

// RegisterHandlersWithOptions creates http.Handler with additional options
func RegisterHandlersWithOptions(router gin.IRouter, si ServerInterface, options GinServerOptions) {
	errorHandler := options.ErrorHandler
	if errorHandler == nil {
		errorHandler = func(c *gin.Context, err error, statusCode int) {
			c.JSON(statusCode, gin.H{"msg": err.Error()})
		}
	}

	wrapper := ServerInterfaceWrapper{
		Handler:            si,
		HandlerMiddlewares: options.Middlewares,
		ErrorHandler:       errorHandler,
	}

	router.GET(options.BaseURL+"/articles", wrapper.GetArticleList)
	router.POST(options.BaseURL+"/articles", wrapper.InitializationArticle)
	router.DELETE(options.BaseURL+"/articles/:uri", wrapper.DeleteArticle)
	router.GET(options.BaseURL+"/articles/:uri", wrapper.GetArticleDetail)
	router.PUT(options.BaseURL+"/articles/:uri/", wrapper.SetArticleVersion)
	router.PATCH(options.BaseURL+"/articles/:uri/category", wrapper.ChangeArticleCategory)
	router.PATCH(options.BaseURL+"/articles/:uri/tags", wrapper.ModifyArticleTags)
	router.GET(options.BaseURL+"/articles/:uri/versions", wrapper.GetArticleVersion)
	router.POST(options.BaseURL+"/articles/:uri/versions", wrapper.CreateNewArticleVersion)
	router.DELETE(options.BaseURL+"/articles/:uri/versions/:version", wrapper.RemoveArticleVersion)
	router.PATCH(options.BaseURL+"/articles/:uri/visibility", wrapper.ChangeArticleVisibility)
	router.GET(options.BaseURL+"/tags", wrapper.GetAllTags)
}
