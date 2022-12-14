// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package auth

import (
	"github.com/YunosukeY/kind-backend/internal/auth/controller"
	"github.com/YunosukeY/kind-backend/internal/auth/repository"
	"github.com/YunosukeY/kind-backend/internal/util"
)

// Injectors from wire.go:

func initializeRouter(service string) (router, func()) {
	tracer, cleanup := util.NewTracer(service)
	client := repository.NewRedis()
	cache := repository.NewCache(tracer, client)
	controllerController := controller.NewController(tracer, cache)
	authRouter := newRouter(controllerController)
	return authRouter, func() {
		cleanup()
	}
}

func initializeDummyRouter(service string) (router, func()) {
	tracer, cleanup := util.NewTracer(service)
	cache := repository.NewDummyCache(tracer)
	controllerController := controller.NewController(tracer, cache)
	authRouter := newRouter(controllerController)
	return authRouter, func() {
		cleanup()
	}
}
