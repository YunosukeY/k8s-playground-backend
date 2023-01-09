package app

import (
	"net"

	"github.com/YunosukeY/kind-backend/internal/grpc"
	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/gin-gonic/gin"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func Serve(dummy bool) {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.With().Caller().Logger()
	gin.SetMode(gin.ReleaseMode)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	s := util.DefaultServer()
	var ts grpc.TodoServiceServer
	var shutdownProvider func()
	if !dummy {
		ts, shutdownProvider = initializeServer("app")
	} else {
		ts, shutdownProvider = initializeDummyServer("app")
	}
	defer shutdownProvider()
	grpc.RegisterTodoServiceServer(s, ts)

	go util.RunPodCommonHandler()
	if err := s.Serve(listener); err != nil {
		log.Panic().Err(err).Msg("")
		panic(err)
	}
}
