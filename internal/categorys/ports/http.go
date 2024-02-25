package ports

import (
	"categorys/application"
	"common/server/httperr"
	"net/http"

	"categorys/application/command"
)

type HttpServer struct {
	app *application.App
}

func NewHttpServer(app *application.App) *HttpServer {
	return &HttpServer{app: app}
}

func (h HttpServer) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var err error
	err = h.app.Commands.CreateCategory.Handle(r.Context(), command.CreateCategory{
		Name:        "blog",
		DisplayName: "探索日志",
		SeoDesc:     "xxx",
	})
	if err != nil {
		httperr.Respond(w, httperr.CommandHandlerErr, err.Error())
		return
	}
	httperr.Success(w)
}
