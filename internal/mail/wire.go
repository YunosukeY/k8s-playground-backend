//go:build wireinject
// +build wireinject

package mail

import (
	"github.com/YunosukeY/kind-backend/internal/mail/repository"
	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/google/wire"
)

func initializeController(service string) (controller, func()) {
	wire.Build(NewController, util.NewTracer, repository.NewQueue, repository.NewReader, repository.NewMailer)
	return controller{}, func() {}
}

func initializeDummyController(service string) (controller, func()) {
	wire.Build(NewController, util.NewTracer, repository.NewDummyQueue, repository.NewDummyMailer)
	return controller{}, func() {}
}
