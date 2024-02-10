package adapter

import (
	"blog/internal/categorys/application"
	"blog/internal/categorys/application/command"
	"encoding/json"
	"net/http"
)

type HttpServer struct {
	app *application.App
}

func NewHttpServer(app *application.App) *HttpServer {
	return &HttpServer{app: app}
}

func (h HttpServer) CreateCategory(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]any)
	defer func() {
		_ = json.NewEncoder(w).Encode(data)
	}()

	err := h.app.CommandsBus.Publish(r.Context(), command.CreateCategory{
		Uid:         0,
		Name:        "",
		DisplayName: "",
		SeoDesc:     "",
	})

	if err != nil {
		return
	}
}
