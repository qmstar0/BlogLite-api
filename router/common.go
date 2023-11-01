package router

import "github.com/gin-gonic/gin"

type Control interface {
	Index(c *gin.Context)
	Create(c *gin.Context)
	Store(c *gin.Context)
	Edit(c *gin.Context)
	Update(c *gin.Context)
	Destroy(c *gin.Context)
}

type Draft interface {
	DraftIndex(c *gin.Context)
	Publish(c *gin.Context)
}

type Trash interface {
	TrashIndex(c *gin.Context)
	UnTrash(c *gin.Context)
}

type Img interface {
	ImgUpload(c *gin.Context)
}

type System interface {
	Index(c *gin.Context)
	Update(c *gin.Context)
}

type Statistics interface {
	Index(c *gin.Context)
}
