package controller

import (
	"github.com/ThreeDotsLabs/watermill/components/cqrs"
)

type Controller struct {
	commandBus *cqrs.CommandBus
}
