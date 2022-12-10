// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package mail

import (
	"github.com/YunosukeY/kind-backend/internal/util"
)

// Injectors from wire.go:

func initializeController(service string) (controller, func()) {
	tracer, cleanup := util.NewTracer(service)
	reader := newReader()
	mailQueue := newQueue(tracer, reader)
	mailMailer := newMailer(tracer)
	mailController := newController(tracer, mailQueue, mailMailer)
	return mailController, func() {
		cleanup()
	}
}

func initializeDummyController(service string) (controller, func()) {
	tracer, cleanup := util.NewTracer(service)
	mailQueue := newDummyQueue(tracer)
	mailMailer := newDummyMailer(tracer)
	mailController := newController(tracer, mailQueue, mailMailer)
	return mailController, func() {
		cleanup()
	}
}
