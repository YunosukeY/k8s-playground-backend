//go:build wireinject
// +build wireinject

package app

import (
	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/google/wire"
)

func initializeRouter(service string) (router, func()) {
	wire.Build(newRouter, newController, util.NewTracer, newRepository, newDB, newQueue, newWriter)
	return router{}, func() {}
}

func initializeDummyRouter(service string) (router, func()) {
	wire.Build(newRouter, newController, util.NewTracer, newDummyRepository, newDummyQueue)
	return router{}, func() {}
}
