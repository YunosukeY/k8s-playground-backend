// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package auth

import (
	"github.com/YunosukeY/kind-backend/internal/util"
)

// Injectors from wire.go:

func initializeRouter(service string) (router, func()) {
	tracer, cleanup := util.NewTracer(service)
	client := newRedis()
	authCache := newCache(tracer, client)
	authController := newController(tracer, authCache)
	authRouter := newRouter(authController)
	return authRouter, func() {
		cleanup()
	}
}
