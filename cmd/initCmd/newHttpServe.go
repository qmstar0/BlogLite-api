package initCmd

import (
	"blog/adapter/httpAdapter"
	"blog/adapter/httpAdapter/httpService"
	"blog/apps/commandResult"
	"github.com/ThreeDotsLabs/watermill/components/requestreply"
)

func NewHttpServe(bus requestreply.CommandBus, backend requestreply.Backend[commandResult.StateCode]) *httpService.HttpServe {
	adapter := httpAdapter.NewHttpAdapter(
		bus,
		backend,
	)
	serve := httpService.NewHttpServe(adapter, httpService.HttpServeConfig{
		Addr: ":3000",
	})
	return serve
}
