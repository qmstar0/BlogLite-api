package main

import (
	categoryAdapter "blog/internal/categorys/adapter"
	"blog/internal/common/server"
	"github.com/go-chi/chi/v5"
	"github.com/qmstar0/eio-cqrs/cqrs"
	"github.com/qmstar0/eio/processor"
	"github.com/qmstar0/eio/pubsub/gopubsub"
)

func main() {
	pubsub := gopubsub.NewGoPubsub("test", gopubsub.GoPubsubConfig{})

	bus := cqrs.NewRouterBus(processor.NewRouter(), cqrs.NewJsonMarshaler(nil))

	categoryApp := categoryAdapter.NewApp(pubsub, pubsub, bus)

	server.RunHttpServer(":3000", func(router chi.Router) {
		router.Route("/api", func(r chi.Router) {
			categoryAdapter.HandlerFromMuxWithBaseURL(categoryAdapter.NewHttpServer(categoryApp), r, "/categorys")
		})
	})
}
