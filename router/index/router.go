package index

import (
	"blog/app/handlers"
	"blog/app/middleware"
	"blog/app/permissions"
	"blog/app/validate"
	"github.com/gin-gonic/gin"
)

// Router 配置路由
func Router() *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	r := router.Group("/api")
	r.Use(middleware.AuthorizerMiddleware())

	V := validate.NewValidate()
	P := permissions.NewPermission()

	handlerArticle := handlers.NewArticle()
	handlerCate := handlers.NewCate()
	handlerTags := handlers.NewTags()
	handlerDraft := handlers.NewDraft()
	handlerTrash := handlers.NewTrash()
	handlerImgUpload := handlers.NewImgUpload()
	handlerUser := handlers.NewUser()
	handlerAuth := handlers.NewAuth()

	a := r.Group("/article")
	{
		artV := V.NewArticleV.Validate()
		authP := P.NewAuthP.Permission()
		a.GET("", handlerArticle.Index)

		a.GET("/create", authP, handlerArticle.Create)
		a.POST("/create", authP, artV, handlerArticle.Store)
		a.GET("/:aid/edit", authP, handlerArticle.Edit)
		a.PUT("/:aid", authP, artV, handlerArticle.Update)
		a.DELETE("/:aid", authP, handlerArticle.Destroy)

		a.GET("/trash", authP, handlerTrash.TrashIndex)
		a.PUT("/:aid/trash", authP, handlerTrash.UnTrash)

		a.GET("/draft", authP, handlerDraft.DraftIndex)
		a.PUT("/:aid/publish", authP, handlerDraft.Publish)

		upload := a.Group("/upload")
		{
			upload.POST("/img", authP, handlerImgUpload.ImgUpload)
		}
	}
	c := r.Group("/cate")
	{
		cateV := V.NewCateV.Validate()
		c.GET("", handlerCate.Index)
		c.POST("/create", cateV, handlerCate.Store)
		c.GET("/:cid/edit", handlerCate.Edit)
		c.PUT("/:cid", cateV, handlerCate.Update)
		c.DELETE("/:cid", handlerCate.Destroy)
		//c.GET("/create", handlerCate.Create)
	}
	t := r.Group("/tags")
	{
		tagV := V.NewTagsV.Validate()
		t.GET("", handlerTags.Index)
		t.POST("/create", tagV, handlerTags.Store)
		t.GET("/:tid/edit", handlerTags.Edit)
		t.PUT("/:tid", tagV, handlerTags.Update)
		t.DELETE("/:tid", handlerTags.Destroy)
		//t.GET("/create", handlerTags.Create)
	}
	u := r.Group("/user")
	captchaV := V.NewCaptchaV.Validate()
	{
		userV := V.NewUserV.Validate()
		reset := u.Group("/reset")
		{
			reset.POST("/pwd", captchaV, handlerUser.ResetPwd)
			reset.GET("/pwd", handlerAuth.AuthResetPwd)
		}
		u.PUT("/update", userV, handlerUser.Update)
	}
	login := r.Group("/login")
	{
		login.POST("", captchaV, handlerUser.Login)
		login.GET("", handlerAuth.AuthLogin)
	}
	register := r.Group("/register")
	{
		register.POST("", captchaV, handlerUser.Register)
		register.GET("", handlerAuth.AuthRegister)
	}
	return router
}
