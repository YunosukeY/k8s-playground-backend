//go:build wireinject
// +build wireinject

package auth

import (
	"github.com/YunosukeY/kind-backend/internal/auth/controller"
	"github.com/YunosukeY/kind-backend/internal/auth/repository"
	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/google/wire"
)

func initializeRouter(service string) (router, func()) {
	wire.Build(newRouter, controller.NewController, util.NewTracer, repository.NewCache, repository.NewRedis)
	return router{}, func() {}
}

func initializeDummyRouter(service string) (router, func()) {
	wire.Build(newRouter, controller.NewController, util.NewTracer, repository.NewDummyCache)
	return router{}, func() {}
}
