package usecase

import (
	"context"

	"github.com/YunosukeY/kind-backend/internal/app/model"
	"github.com/YunosukeY/kind-backend/internal/app/repository"
	"github.com/YunosukeY/kind-backend/internal/util"
	"go.opentelemetry.io/otel/trace"
)

type Usecase interface {
	GetAllTodos(context.Context) ([]model.TodoForResponse, error)
	CreateTodo(context.Context, model.TodoForPostRequest) (*model.TodoForResponse, error)
	SendMail(context.Context, model.Mail) error
}

type usecase struct {
	t trace.Tracer
	r repository.Repository
	q repository.Queue
}

func NewUsecase(t trace.Tracer, r repository.Repository, q repository.Queue) Usecase {
	return usecase{t, r, q}
}

func (u usecase) GetAllTodos(ctx context.Context) ([]model.TodoForResponse, error) {
	child, span := u.t.Start(ctx, util.FuncName())
	defer span.End()

	return u.r.FindAllTodos(child)
}

func (u usecase) CreateTodo(ctx context.Context, todo model.TodoForPostRequest) (*model.TodoForResponse, error) {
	child, span := u.t.Start(ctx, util.FuncName())
	defer span.End()

	return u.r.CreateTodo(child, todo)
}

func (u usecase) SendMail(ctx context.Context, mail model.Mail) error {
	child, span := u.t.Start(ctx, util.FuncName())
	defer span.End()

	return u.q.Push(child, mail)
}
