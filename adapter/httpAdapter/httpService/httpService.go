package httpService

import "github.com/go-chi/chi/v5"

type HttpServe struct {
	Addr string
	Mux  *chi.Mux
}

func NewHttpServeWithDefault() *HttpServe {
	return &HttpServe{
		Addr: ":3000",
		Mux:  chi.NewRouter(),
	}
}

func NewHttpServe(addr string, mux *chi.Mux) *HttpServe {
	return &HttpServe{
		Addr: addr,
		Mux:  mux,
	}
}
