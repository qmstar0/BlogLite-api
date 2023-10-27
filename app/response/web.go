package response

import "github.com/gin-gonic/gin"

// Web html响应
type Web struct {
	C *gin.Context
}

//func (w Web) Success(path string) {
//	w.C.HTML(http.StatusOK, path)
//	w.C.Abort()
//}
