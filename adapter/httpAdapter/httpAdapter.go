package httpAdapter

import (
	"blog/adapter/httpAdapter/router"
	"context"
	"fmt"
	"github.com/qmstar0/eio-cqrs/cqrs"
	"github.com/qmstar0/eio/message"
	"net/http"
	"time"
)

type HttpAdapter struct {
	bus cqrs.PublishBus
}

func NewHttpAdapter(bus cqrs.Bus) *HttpAdapter {
	return &HttpAdapter{bus: bus}
}

func (h HttpAdapter) Adapter(constructor router.CommandConstructor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cmd, err := constructor(w, r)
		if err != nil {
			w.WriteHeader(500)
			return
		}

		ctx, cancelFunc := context.WithTimeout(context.WithValue(r.Context(), httpAdapterCtxKey, w), time.Second)
		defer cancelFunc()

		if err != nil {
			w.WriteHeader(500)
			return
		}
	}
}

var httpAdapterCtxKey = &ctxKey{"httpAdapterCtxKey"}

type ctxKey struct {
	name string
}

func (c ctxKey) String() string {
	return fmt.Sprintf("HttpAdapter ctxKey `%s`", c.name)
}

func RespondCallback(msg *message.Context) {
	respW := msg.Value(httpAdapterCtxKey).(http.ResponseWriter)
	err := msg.Err()
}
