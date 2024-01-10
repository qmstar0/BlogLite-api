package router

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type CommandConstructor func(r *http.Request) (cmd any, err error)

type HttpHandleConstructor interface {
	HttpHandle(CommandConstructor) http.HandlerFunc
}

func SetUpRoutes(mux *chi.Mux, constructor HttpHandleConstructor) {
	// -/api/*
	mux.Route("/api", func(api chi.Router) {

		// -/api/category/*
		api.Route("/category", func(category chi.Router) {

			category.Post("/create", constructor.HttpHandle(CreateCategroy))

		})
	})
}
