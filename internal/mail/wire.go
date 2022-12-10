//go:build wireinject
// +build wireinject

package mail

import (
	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/google/wire"
)

func initializeController(service string) (controller, func()) {
	wire.Build(newController, util.NewTracer, newQueue, newReader, newMailer)
	return controller{}, func() {}
}
