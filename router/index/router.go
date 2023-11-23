package index

import (
	"blog/app/handlers"
	m "blog/app/middleware"
	"blog/app/validate"
	"github.com/gin-gonic/gin"
)

// Router 配置路由
func Router() *gin.Engine {
	router := gin.Default()
	router.Use(m.CORSMiddleware())

	r := router.Group("/api")
	r.Use(m.AuthorizerMiddleware())
	V := validate.NewValidate()

	handlerArticle := handlers.NewArticle()
	handlerCate := handlers.NewCate()
	handlerTags := handlers.NewTags()
	handlerStatus := handlers.NewStatus()
	handlerImgUpload := handlers.NewImgUpload()
	handlerUser := handlers.NewUser()
	handlerAuth := handlers.NewAuth()

	a := r.Group("/article")
	{
		artV := V.NewArticleV.Validate()
		a.GET("", handlerArticle.Index)
		pa := a.Group("")
		{
			pa.GET("/create", handlerArticle.Create)
			pa.POST("/create", artV, handlerArticle.Store)
			pa.GET("/:aid/edit", handlerArticle.Edit)
			pa.PUT("/:aid", artV, handlerArticle.Update)
			pa.DELETE("/:aid", handlerArticle.Destroy)

			pa.GET("/trash", handlerStatus.TrashIndex)
			pa.GET("/draft", handlerStatus.DraftIndex)
			pa.PUT("/:aid/publish", handlerStatus.Publish)
		}
		upload := a.Group("/upload")
		{
			upload.POST("/img", handlerImgUpload.ImgUpload)
		}
	}
	c := r.Group("/cate")
	{
		cateV := V.NewCateV.Validate()
		c.GET("", handlerCate.Index)
		pc := c.Group("")
		{
			pc.POST("/create", cateV, handlerCate.Store)
			pc.GET("/:cid/edit", handlerCate.Edit)
			pc.PUT("/:cid", cateV, handlerCate.Update)
			pc.DELETE("/:cid", handlerCate.Destroy)
		}
		//c.GET("/create", handlerCate.Create)
	}
	t := r.Group("/tags")
	{
		tagV := V.NewTagsV.Validate()
		t.GET("", handlerTags.Index)
		pt := t.Group("").Use()
		{
			pt.POST("/create", tagV, handlerTags.Store)
			pt.GET("/:tid/edit", handlerTags.Edit)
			pt.PUT("/:tid", tagV, handlerTags.Update)
			pt.DELETE("/:tid", handlerTags.Destroy)
		}
		//t.GET("/create", handlerTags.Create)
	}
	u := router.Group("/user")
	captchaV := V.NewCaptchaV.Validate()
	{
		u.POST("/login", captchaV, handlerUser.Login)
		u.POST("/register", captchaV, handlerUser.Register)
	}
	u2 := r.Group("/user")
	{
		userV := V.NewUserV.Validate()
		u2.PUT("/update", userV, handlerUser.Update)
		u2.PUT("/reset/pwd", captchaV, handlerUser.ResetPwd)
	}
	auth := router.Group("auth")
	{
		auth.GET("/login", handlerAuth.AuthLogin)
		auth.GET("/register", handlerAuth.AuthRegister)
		auth.Use(m.AuthorizerMiddleware()).GET("/reset/pwd", handlerAuth.AuthResetPwd)
	}
	return router
}
