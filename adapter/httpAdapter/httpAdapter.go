package httpAdapter

import (
	"blog/adapter/httpAdapter/httpService"
	"blog/adapter/httpAdapter/respond"
	"blog/apps/commandResult"
	"context"
	"github.com/ThreeDotsLabs/watermill/components/requestreply"
	"net/http"
)

type CommandConstructor func(r *http.Request) (cmd any, err error)

type HttpAdapter struct {
	bus     requestreply.CommandBus
	backend requestreply.Backend[commandResult.StateCode]
	Serve   *httpService.HttpServe
}

func NewHttpAdapter(bus requestreply.CommandBus, backend requestreply.Backend[commandResult.StateCode]) *HttpAdapter {
	return &HttpAdapter{bus: bus, backend: backend}
}

func (h HttpAdapter) Handle(fn CommandConstructor) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			err       error
			stateCode commandResult.StateCode
		)
		defer respond.Respond(w, &stateCode)

		cmd, err := fn(r)
		if err != nil {
			stateCode = commandResult.InvalidParam
			return
		}

		reply, err := requestreply.SendWithReply(context.Background(), h.bus, h.backend, &cmd)
		if err != nil {
			stateCode = commandResult.CommandProcessErr
			return
		}
		stateCode = reply.HandlerResult
	}
}

func (h HttpAdapter) StartRun() error {
	return http.ListenAndServe(h.Serve.Addr, h.Serve.Mux)
}

//func (c HttpAdapter) CreateCategory(w http.ResponseWriter, r *http.Request) {
//	var (
//		err       error
//		stateCode commandResult.StateCode
//		category  categorys.CreateCategoryCommand
//	)
//	defer respond.Respond(w, &stateCode)
//
//	fmt.Println(r.Form["key"])
//
//	fmt.Println(r.FormValue("key"))
//
//	body, err := io.ReadAll(r.Body)
//	if err != nil {
//		stateCode = commandResult.InvalidParam
//		return
//	}
//
//	err = json.Unmarshal(body, &category)
//	if err != nil {
//		stateCode = commandResult.InvalidParam
//		return
//	}
//
//	ctx := context.Background()
//	reply, err := requestreply.SendWithReply(ctx, c.bus, c.backend, &category)
//	if err != nil {
//		stateCode = commandResult.CommandProcessErr
//		return
//	}
//	stateCode = reply.HandlerResult
//	stateCode = 0
//}
//
//func (c HttpAdapter) DeleteCategory(w http.ResponseWriter, r *http.Request) {
//	var (
//		err       error
//		stateCode commandResult.StateCode
//		category  categorys.DeleteCategoryCommand
//	)
//	defer respond.Respond(w, &stateCode)
//
//	body, err := io.ReadAll(r.Body)
//	if err != nil {
//		stateCode = commandResult.InvalidParam
//		return
//	}
//
//	err = json.Unmarshal(body, &category)
//	if err != nil {
//		stateCode = commandResult.InvalidParam
//		return
//	}
//
//	ctx := context.Background()
//	reply, err := requestreply.SendWithReply(ctx, c.bus, c.backend, &category)
//	if err != nil {
//		stateCode = commandResult.CommandProcessErr
//		return
//	}
//	stateCode = reply.HandlerResult
//}
//
//func (c HttpAdapter) UpdateCategory(w http.ResponseWriter, r *http.Request) {
//	var (
//		err       error
//		stateCode commandResult.StateCode
//		category  categorys.UpdateCategoryCommand
//	)
//	defer respond.Respond(w, &stateCode)
//	body, err := io.ReadAll(r.Body)
//	if err != nil {
//		stateCode = commandResult.InvalidParam
//		return
//	}
//
//	err = json.Unmarshal(body, &category)
//	if err != nil {
//		stateCode = commandResult.InvalidParam
//		return
//	}
//
//	ctx := context.Background()
//	reply, err := requestreply.SendWithReply(ctx, c.bus, c.backend, &category)
//	if err != nil {
//		stateCode = commandResult.CommandProcessErr
//		return
//	}
//	stateCode = reply.HandlerResult
//}
