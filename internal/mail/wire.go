//go:build wireinject
// +build wireinject

package mail

import (
	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/google/wire"
)

func initializeController(service string) (controller, func()) {
	wire.Build(NewController, util.NewTracer, NewQueue, NewReader, NewMailer)
	return controller{}, func() {}
}

func initializeDummyController(service string) (controller, func()) {
	wire.Build(NewController, util.NewTracer, NewDummyQueue, NewDummyMailer)
	return controller{}, func() {}
}
