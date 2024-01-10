package httpAdapter

import (
	"blog/adapter/httpAdapter/httpService"
	"blog/adapter/httpAdapter/respond"
	"blog/adapter/httpAdapter/router"
	"blog/apps/commandResult"
	"context"
	"github.com/ThreeDotsLabs/watermill/components/requestreply"
	"net/http"
)

type HttpAdapter struct {
	bus     requestreply.CommandBus
	backend requestreply.Backend[commandResult.StateCode]
	Serve   *httpService.HttpServe
}

func NewHttpAdapter(
	bus requestreply.CommandBus,
	backend requestreply.Backend[commandResult.StateCode],
	serve *httpService.HttpServe,
) *HttpAdapter {
	adapter := &HttpAdapter{bus: bus, backend: backend, Serve: serve}
	router.SetUpRoutes(adapter.Serve.Mux, adapter)
	return adapter
}

func (h HttpAdapter) HttpHandle(fn router.CommandConstructor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			err error
		)
		defer respond.Respond(w, &err)

		cmd, err := fn(r)
		if err != nil {
			return
		}

		reply, err := requestreply.SendWithReply(context.Background(), h.bus, h.backend, &cmd)
		if err != nil {
			return
		}
		err = respond.NewError(reply.HandlerResult)
	}
}

func (h HttpAdapter) StartRun() error {
	return http.ListenAndServe(h.Serve.Addr, h.Serve.Mux)
}
