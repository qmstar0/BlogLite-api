// Package ports provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen/v2 version v2.1.0 DO NOT EDIT.
package ports

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oapi-codegen/runtime"
)

// CreateCategoryJSONBody defines parameters for CreateCategory.
type CreateCategoryJSONBody struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Slug        string `json:"slug"`
}

// ModifyCategoryDescriptionJSONBody defines parameters for ModifyCategoryDescription.
type ModifyCategoryDescriptionJSONBody struct {
	Descripition string `json:"descripition"`
}

// CreateCategoryJSONRequestBody defines body for CreateCategory for application/json ContentType.
type CreateCategoryJSONRequestBody CreateCategoryJSONBody

// ModifyCategoryDescriptionJSONRequestBody defines body for ModifyCategoryDescription for application/json ContentType.
type ModifyCategoryDescriptionJSONRequestBody ModifyCategoryDescriptionJSONBody

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// 获取分类列表
	// (GET /categories)
	GetCategoryList(c *gin.Context)
	// 创建分类
	// (POST /categories)
	CreateCategory(c *gin.Context)
	// 删除分类
	// (DELETE /categories/{slug})
	DeleteCategory(c *gin.Context, slug string)
	// 更改分类描述
	// (PATCH /categories/{slug})
	ModifyCategoryDescription(c *gin.Context, slug string)
}

// ServerInterfaceWrapper converts contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler            ServerInterface
	HandlerMiddlewares []MiddlewareFunc
	ErrorHandler       func(*gin.Context, error, int)
}

type MiddlewareFunc func(c *gin.Context)

// GetCategoryList operation middleware
func (siw *ServerInterfaceWrapper) GetCategoryList(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.GetCategoryList(c)
}

// CreateCategory operation middleware
func (siw *ServerInterfaceWrapper) CreateCategory(c *gin.Context) {

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.CreateCategory(c)
}

// DeleteCategory operation middleware
func (siw *ServerInterfaceWrapper) DeleteCategory(c *gin.Context) {

	var err error

	// ------------- Path parameter "slug" -------------
	var slug string

	err = runtime.BindStyledParameterWithOptions("simple", "slug", c.Param("slug"), &slug, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter slug: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.DeleteCategory(c, slug)
}

// ModifyCategoryDescription operation middleware
func (siw *ServerInterfaceWrapper) ModifyCategoryDescription(c *gin.Context) {

	var err error

	// ------------- Path parameter "slug" -------------
	var slug string

	err = runtime.BindStyledParameterWithOptions("simple", "slug", c.Param("slug"), &slug, runtime.BindStyledParameterOptions{Explode: false, Required: true})
	if err != nil {
		siw.ErrorHandler(c, fmt.Errorf("Invalid format for parameter slug: %w", err), http.StatusBadRequest)
		return
	}

	for _, middleware := range siw.HandlerMiddlewares {
		middleware(c)
		if c.IsAborted() {
			return
		}
	}

	siw.Handler.ModifyCategoryDescription(c, slug)
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

	router.GET(options.BaseURL+"/categories", wrapper.GetCategoryList)
	router.POST(options.BaseURL+"/categories", wrapper.CreateCategory)
	router.DELETE(options.BaseURL+"/categories/:slug", wrapper.DeleteCategory)
	router.PATCH(options.BaseURL+"/categories/:slug", wrapper.ModifyCategoryDescription)
}
