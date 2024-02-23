package ports

import (
	"categorys/application"
	"categorys/application/command"
	"common/httperr"
	"net/http"
)

type HttpServer struct {
	app *application.App
}

func NewHttpServer(app *application.App) *HttpServer {
	return &HttpServer{app: app}
}

func (h HttpServer) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var err error
	defer httperr.Respond(w, &err)
	err = h.app.Commands.CreateCategory.Handle(r.Context(), command.CreateCategory{
		Name:        "blog",
		DisplayName: "探索日志",
		SeoDesc:     "xxx",
	})
	if err != nil {
		return
	}
}
