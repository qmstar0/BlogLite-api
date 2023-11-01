package validate

import (
	"blog/app/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var (
	validate = validator.New()
	Srv      = service.GetSrv()
)

type V interface {
	Validate() gin.HandlerFunc
}

type Validate struct {
	NewArticleV V
	NewCateV    V
	NewTagsV    V
	NewUserV    V
	NewAuthV    V
	NewCaptchaV V
}

func NewValidate() *Validate {
	return &Validate{
		NewArticleV: ArticleV{},
		NewCateV:    CateV{},
		NewTagsV:    TagV{},
		NewUserV:    UserV{},
		NewAuthV:    AuthV{},
		NewCaptchaV: CaptchaV{},
	}
}
