package controller

import (
	"context"

	"github.com/YunosukeY/kind-backend/internal/app/model"
	"github.com/YunosukeY/kind-backend/internal/app/usecase"
	"github.com/YunosukeY/kind-backend/internal/grpc"
	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
)

type Server struct {
	t trace.Tracer
	grpc.UnimplementedTodoServiceServer
	u usecase.Usecase
}

func NewServer(t trace.Tracer, u usecase.Usecase) grpc.TodoServiceServer {
	return Server{t: t, u: u}
}

func (c Server) ListTodos(ctx context.Context, _ *grpc.ListTodosRequest) (*grpc.ListTodosResponse, error) {
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

	return &grpc.ListTodosResponse{Todos: todos}, nil
}

func (c Server) CreateTodo(ctx context.Context, req *grpc.CreateTodoRequest) (*grpc.Todo, error) {
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
