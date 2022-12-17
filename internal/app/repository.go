package app

import (
	"context"

	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/rs/zerolog/log"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"go.opentelemetry.io/otel/trace"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewDB() *gorm.DB {
	dsn := util.GetConnectionString()

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Panic().Err(err).Msg("")
		panic(err)
	}
	if err := db.Use(otelgorm.NewPlugin()); err != nil {
		log.Panic().Err(err).Msg("")
		panic(err)
	}
	return db
}

type Repository interface {
	findAllTodos(ctx context.Context) ([]TodoForResponse, error)
	createTodo(ctx context.Context, todo TodoForPostRequest) (*TodoForResponse, error)
}

type repository struct {
	t  trace.Tracer
	db *gorm.DB
}

func NewRepository(t trace.Tracer, db *gorm.DB) Repository {
	return repository{t, db}
}

func (r repository) findAllTodos(ctx context.Context) ([]TodoForResponse, error) {
	child, span := r.t.Start(ctx, util.FuncName())
	defer span.End()

	var todos []TodoForResponse
	res := r.db.WithContext(child).Table("todos").Find(&todos)
	if res.Error != nil {
		log.Error().Err(res.Error).Msg("")
		return nil, res.Error
	}
	return todos, nil
}

func (r repository) createTodo(ctx context.Context, todo TodoForPostRequest) (*TodoForResponse, error) {
	child, span := r.t.Start(ctx, util.FuncName())
	defer span.End()

	todoWithID := &TodoForResponse{Content: todo.Content}
	res := r.db.WithContext(child).Table("todos").Create(todoWithID)
	if res.Error != nil {
		log.Error().Err(res.Error).Msg("")
		return nil, res.Error
	}
	return todoWithID, nil
}

type dummyRepository struct {
	t         trace.Tracer
	totalSize int
	todos     []TodoForResponse
}

func NewDummyRepository(t trace.Tracer) Repository {
	return &dummyRepository{t: t, totalSize: 0, todos: []TodoForResponse{}}
}

func (r *dummyRepository) findAllTodos(ctx context.Context) ([]TodoForResponse, error) {
	_, span := r.t.Start(ctx, util.FuncName())
	defer span.End()

	return r.todos, nil
}

func (r *dummyRepository) createTodo(ctx context.Context, todo TodoForPostRequest) (*TodoForResponse, error) {
	_, span := r.t.Start(ctx, util.FuncName())
	defer span.End()

	todoWithID := TodoForResponse{ID: r.totalSize, Content: todo.Content}
	r.todos = append(r.todos, todoWithID)
	r.totalSize++

	return &todoWithID, nil
}
