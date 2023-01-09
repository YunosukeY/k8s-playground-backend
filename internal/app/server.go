package app

import (
	"context"
	"net"

	"github.com/YunosukeY/kind-backend/internal/app/model"
	"github.com/YunosukeY/kind-backend/internal/app/usecase"
	"github.com/YunosukeY/kind-backend/internal/grpc"
	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/gin-gonic/gin"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type server struct {
	t trace.Tracer
	grpc.UnimplementedTodoServiceServer
	u usecase.Usecase
}

func newServer(t trace.Tracer, u usecase.Usecase) grpc.TodoServiceServer {
	return server{t: t, u: u}
}

func (c server) GetAll(ctx context.Context, _ *emptypb.Empty) (*grpc.Todos, error) {
	child, span := c.t.Start(ctx, util.FuncName())
	defer span.End()

	ts, err := c.u.GetAllTodos(child)
	if err != nil {
		return nil, err
	}
	log.Debug().Interface("todos", ts).Msg("")

	todos := []*grpc.Todo{}
	for _, t := range ts {
		todo := grpc.Todo{Id: int32(t.ID), Content: t.Content}
		todos = append(todos, &todo)
	}

	return &grpc.Todos{Todos: todos}, nil
}

func (c server) Create(ctx context.Context, req *grpc.CreateTodoRequest) (*grpc.Todo, error) {
	child, span := c.t.Start(ctx, util.FuncName())
	defer span.End()

	todo := model.TodoForPostRequest{Content: req.GetContent()}

	todoWithID, err := c.u.CreateTodo(child, todo)
	if err != nil {
		return nil, err
	}
	log.Debug().Interface("todo", todoWithID).Msg("")

	return &grpc.Todo{Id: int32(todoWithID.ID), Content: todoWithID.Content}, nil
}

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
