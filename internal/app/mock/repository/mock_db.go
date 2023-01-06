// Code generated by MockGen. DO NOT EDIT.
// Source: internal/app/repository/db.go

// Package mock_repository is a generated GoMock package.
package mock_repository

import (
	context "context"
	reflect "reflect"

	model "github.com/YunosukeY/kind-backend/internal/app/model"
	gomock "github.com/golang/mock/gomock"
)

// MockRepository is a mock of Repository interface.
type MockRepository struct {
	ctrl     *gomock.Controller
	recorder *MockRepositoryMockRecorder
}

// MockRepositoryMockRecorder is the mock recorder for MockRepository.
type MockRepositoryMockRecorder struct {
	mock *MockRepository
}

// NewMockRepository creates a new mock instance.
func NewMockRepository(ctrl *gomock.Controller) *MockRepository {
	mock := &MockRepository{ctrl: ctrl}
	mock.recorder = &MockRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepository) EXPECT() *MockRepositoryMockRecorder {
	return m.recorder
}

// CreateTodo mocks base method.
func (m *MockRepository) CreateTodo(ctx context.Context, todo model.TodoForPostRequest) (*model.TodoForResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateTodo", ctx, todo)
	ret0, _ := ret[0].(*model.TodoForResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateTodo indicates an expected call of CreateTodo.
func (mr *MockRepositoryMockRecorder) CreateTodo(ctx, todo interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateTodo", reflect.TypeOf((*MockRepository)(nil).CreateTodo), ctx, todo)
}

// FindAllTodos mocks base method.
func (m *MockRepository) FindAllTodos(ctx context.Context) ([]model.TodoForResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAllTodos", ctx)
	ret0, _ := ret[0].([]model.TodoForResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindAllTodos indicates an expected call of FindAllTodos.
func (mr *MockRepositoryMockRecorder) FindAllTodos(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAllTodos", reflect.TypeOf((*MockRepository)(nil).FindAllTodos), ctx)
}
