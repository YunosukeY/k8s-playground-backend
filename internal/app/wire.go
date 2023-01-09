//go:build wireinject
// +build wireinject

package app

import (
	"github.com/YunosukeY/kind-backend/internal/app/controller"
	"github.com/YunosukeY/kind-backend/internal/app/repository"
	"github.com/YunosukeY/kind-backend/internal/app/usecase"
	"github.com/YunosukeY/kind-backend/internal/grpc"
	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/google/wire"
)

func initializeRouter(service string) (router, func()) {
	wire.Build(newRouter, controller.NewRestController, util.NewTracer, usecase.NewUsecase, repository.NewRepository, repository.NewDB, repository.NewQueue, repository.NewWriter)
	return router{}, func() {}
}

func initializeDummyRouter(service string) (router, func()) {
	wire.Build(newRouter, controller.NewRestController, util.NewTracer, usecase.NewUsecase, repository.NewDummyRepository, repository.NewDummyQueue)
	return router{}, func() {}
}

func initializeServer(service string) (grpc.TodoServiceServer, func()) {
	wire.Build(newServer, util.NewTracer, usecase.NewUsecase, repository.NewRepository, repository.NewDB, repository.NewQueue, repository.NewWriter)
	return server{}, func() {}
}

func initializeDummyServer(service string) (grpc.TodoServiceServer, func()) {
	wire.Build(newServer, util.NewTracer, usecase.NewUsecase, repository.NewDummyRepository, repository.NewDummyQueue)
	return server{}, func() {}
}
