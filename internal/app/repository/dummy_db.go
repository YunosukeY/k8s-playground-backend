package repository

import (
	"context"

	"github.com/YunosukeY/kind-backend/internal/app/model"
	"github.com/YunosukeY/kind-backend/internal/util"
	"go.opentelemetry.io/otel/trace"
)

type dummyRepository struct {
	t         trace.Tracer
	totalSize int
	todos     []model.TodoForResponse
}

func NewDummyRepository(t trace.Tracer) Repository {
	return &dummyRepository{t: t, totalSize: 0, todos: []model.TodoForResponse{}}
}

func (r *dummyRepository) FindAllTodos(ctx context.Context) ([]model.TodoForResponse, error) {
	_, span := r.t.Start(ctx, util.FuncName())
	defer span.End()

	return r.todos, nil
}

func (r *dummyRepository) CreateTodo(ctx context.Context, todo model.TodoForPostRequest) (*model.TodoForResponse, error) {
	_, span := r.t.Start(ctx, util.FuncName())
	defer span.End()

	todoWithID := model.TodoForResponse{ID: r.totalSize, Content: todo.Content}
	r.todos = append(r.todos, todoWithID)
	r.totalSize++

	return &todoWithID, nil
}
