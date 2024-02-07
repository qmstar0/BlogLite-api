package router

import (
	"blog/adapter/httpAdapter/router/w"
	"github.com/go-chi/chi/v5"
	"net/http"
)

type CommandConstructor func(w http.ResponseWriter, r *http.Request) (cmd any, err error)

type HttpHandle interface {
	Adapter(CommandConstructor) http.HandlerFunc
}

func SetUpRoutes(mux *chi.Mux, adapter HttpHandle) {
	// -/api/*
	mux.Route("/api", func(api chi.Router) {

		// -/api/category/*
		api.Route("/category", func(category chi.Router) {

			category.Post("/create", adapter.Adapter(w.CreateCategroy))

		})
	})
}
