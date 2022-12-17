//go:build wireinject
// +build wireinject

package app

import (
	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/google/wire"
)

func initializeRouter(service string) (router, func()) {
	wire.Build(newRouter, NewController, util.NewTracer, NewRepository, NewDB, NewQueue, NewWriter)
	return router{}, func() {}
}

func initializeDummyRouter(service string) (router, func()) {
	wire.Build(newRouter, NewController, util.NewTracer, NewDummyRepository, NewDummyQueue)
	return router{}, func() {}
}
