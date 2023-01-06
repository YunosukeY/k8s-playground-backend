package repository

import (
	"context"
	"testing"

	"github.com/YunosukeY/kind-backend/internal/app/model"
	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/stretchr/testify/assert"
)

func TestFindAllTodos(t *testing.T) {
	db, deferer := util.NewTestDB("find_all_todos")
	defer deferer()
	tracer, shutdownProvider := util.NewTracer("app")
	defer shutdownProvider()
	repo := NewRepository(tracer, db)

	expected := []model.TodoForResponse{{ID: 1, Content: "test"}}
	actual, _ := repo.FindAllTodos(context.Background())
	assert.Equal(t, expected, actual)
}

func TestCreateTodo(t *testing.T) {
	db, deferer := util.NewTestDB("create_todo")
	defer deferer()
	tracer, shutdownProvider := util.NewTracer("app")
	defer shutdownProvider()
	repo := NewRepository(tracer, db)

	expected := &model.TodoForResponse{ID: 2, Content: "todo2"}
	input := model.TodoForPostRequest{Content: "todo2"}
	actual, _ := repo.CreateTodo(context.Background(), input)
	assert.Equal(t, expected.Content, actual.Content) // txdbでロールバックしてもauto_incrementなIDは戻らないのでcontentのみ比較
	todos, _ := repo.FindAllTodos(context.Background())
	assert.Len(t, todos, 2)
	assert.Equal(t, 2, len(todos))
}
