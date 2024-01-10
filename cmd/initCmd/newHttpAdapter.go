package initCmd

import (
	"blog/adapter/httpAdapter"
	"blog/adapter/httpAdapter/httpService"
	"blog/apps/commandResult"
	"github.com/ThreeDotsLabs/watermill/components/requestreply"
)

func NewHttpAdapter(bus requestreply.CommandBus, backend requestreply.Backend[commandResult.StateCode]) *httpAdapter.HttpAdapter {
	serve := httpService.NewHttpServeWithDefault()
	return httpAdapter.NewHttpAdapter(
		bus,
		backend,
		serve,
	)
}
