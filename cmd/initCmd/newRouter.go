package initCmd

import (
	"blog/adapter/httpAdapter"
	"blog/apps/commandConstructor"
	"github.com/go-chi/chi/v5"
)

type DTO struct {
	A string
	B int
}

func NewRouter() *httpAdapter.HttpAdapter {
	adapter := httpAdapter.NewHttpAdapter()
	router := chi.NewRouter()

	router.Post("/", adapter.Handle(commandConstructor.CreateCategory))
	return adapter
}
