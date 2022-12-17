//go:build wireinject
// +build wireinject

package auth

import (
	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/google/wire"
)

func initializeRouter(service string) (router, func()) {
	wire.Build(newRouter, NewController, util.NewTracer, NewCache, NewRedis)
	return router{}, func() {}
}

func initializeDummyRouter(service string) (router, func()) {
	wire.Build(newRouter, NewController, util.NewTracer, NewDummyCache)
	return router{}, func() {}
}
