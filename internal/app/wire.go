//go:build wireinject
// +build wireinject

package app

import (
	"github.com/YunosukeY/kind-backend/internal/app/controller"
	"github.com/YunosukeY/kind-backend/internal/app/repository"
	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/google/wire"
)

func initializeRouter(service string) (router, func()) {
	wire.Build(newRouter, controller.NewController, util.NewTracer, repository.NewRepository, repository.NewDB, repository.NewQueue, repository.NewWriter)
	return router{}, func() {}
}

func initializeDummyRouter(service string) (router, func()) {
	wire.Build(newRouter, controller.NewController, util.NewTracer, repository.NewDummyRepository, repository.NewDummyQueue)
	return router{}, func() {}
}
