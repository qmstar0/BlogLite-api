package ports

import (
	"categorys/application"
	"categorys/application/command"
	"categorys/application/query"
	"common/e"
	"common/server/binder"
	"common/server/httperr"
	"net/http"
)

type HttpServer struct {
	app    *application.App
	Binder binder.Binder
}

func NewHttpServer(app *application.App) *HttpServer {
	return &HttpServer{app: app, Binder: binder.JSONBinder{}}
}

func (h HttpServer) GetAllCategory(w http.ResponseWriter, r *http.Request, params GetAllCategoryParams) {
	httperr.Success(w)
}

func (h HttpServer) CreateCategory(w http.ResponseWriter, r *http.Request) {
	var req CreateCategoryJSONRequestBody
	err := h.Binder.Bind(r, &req)
	if err != nil || req.Name == "" || req.DisplayName == "" {
		httperr.Error(w, e.Wrap(e.InvalidParam, err))
		return
	}

	err = h.app.Commands.CreateCategory.Handle(r.Context(), command.CreateCategory{
		Name:        req.Name,
		DisplayName: req.DisplayName,
		SeoDesc:     req.SeoDesc,
	})
	if err != nil {
		httperr.Error(w, e.Wrap(e.CommandHandlerErr, err))
		return
	}

	w.Header().Set("Location", r.URL.JoinPath(req.Name).String())
	httperr.Success(w)
}

func (h HttpServer) DeleteCategory(w http.ResponseWriter, r *http.Request, name string) {
	err := h.app.Commands.DeleteCategory.Handle(r.Context(), command.DeleteCategory{Name: name})
	if err != nil {
		httperr.Error(w, e.Wrap(e.CommandHandlerErr, err))
		return
	}
	httperr.Success(w)
}

func (h HttpServer) GetCategory(w http.ResponseWriter, r *http.Request, name string) {
	result, err := h.app.Queries.GetCategory.Handle(r.Context(), query.GetCategory{Name: name})
	if err != nil {
		r.Context()
		httperr.Error(w, e.Wrap(e.QueryHandlerErr, err))
		return
	}
	httperr.Respond(w, result)
}

func (h HttpServer) UpdateCategory(w http.ResponseWriter, r *http.Request, name string) {
	var req UpdateCategoryJSONRequestBody
	err := h.Binder.Bind(r, &req)
	if err != nil {
		httperr.Error(w, e.Wrap(e.InvalidParam, err))
		return
	}

	if req.SeoDesc == "" && req.DisplayName == "" {
		httperr.Error(w, e.Wrap(e.InvalidParam, err))
		return
	}

	err = h.app.Commands.UpdateCategory.Handle(r.Context(), command.UpdateCategory{
		Name:        name,
		DisplayName: req.DisplayName,
		SeoDesc:     req.SeoDesc,
	})
	if err != nil {
		httperr.Error(w, e.Wrap(e.CommandHandlerErr, err))
		return
	}

	httperr.Success(w)
}
