package main

import (
	"context"
	"testing"

	"github.com/YunosukeY/kind-backend/internal/grpc"
	"github.com/stretchr/testify/assert"
	grpclib "google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestGrpc(t *testing.T) {
	conn, err := grpclib.Dial(
		"localhost:8080",
		grpclib.WithTransportCredentials(insecure.NewCredentials()),
		grpclib.WithBlock(),
	)
	assert.Nil(t, err)
	defer conn.Close()
	client := grpc.NewTodoServiceClient(conn)

	todos, err := client.GetAll(context.Background(), &emptypb.Empty{})
	assert.Nil(t, err)
	assert.Len(t, todos.GetTodos(), 0)

	todo, err := client.Create(context.Background(), &grpc.CreateTodoRequest{Content: "test"})
	assert.Nil(t, err)
	assert.Equal(t, "test", todo.GetContent())

	todos, err = client.GetAll(context.Background(), &emptypb.Empty{})
	assert.Nil(t, err)
	assert.Len(t, todos.GetTodos(), 1)
}
