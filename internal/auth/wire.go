//go:build wireinject
// +build wireinject

package auth

import (
	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/google/wire"
)

func initializeRouter(service string) (router, func()) {
	wire.Build(newRouter, newController, util.NewTracer, newCache, newRedis)
	return router{}, func() {}
}

func initializeDummyRouter(service string) (router, func()) {
	wire.Build(newRouter, newController, util.NewTracer, newDummyCache)
	return router{}, func() {}
}
