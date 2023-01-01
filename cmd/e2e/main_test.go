package main

import (
	"fmt"
	"testing"
	"time"

	appmodel "github.com/YunosukeY/kind-backend/internal/app/model"
	authmodel "github.com/YunosukeY/kind-backend/internal/auth/model"
	"github.com/YunosukeY/kind-backend/internal/e2e"
	"github.com/stretchr/testify/assert"
)

var login = e2e.NewPostAPI[authmodel.User, interface{}]("localhost:8081", "/api/v1/sessions")
var checkLogin = e2e.NewGetAPI[interface{}]("localhost:8081", "/api/v1/sessions")

var publicGetTodos = e2e.NewGetAPI[[]appmodel.TodoForResponse]("localhost:8080", "/api/v1/public/todos")
var getTodos = e2e.NewGetAPI[[]appmodel.TodoForResponse]("localhost:8080", "/api/v1/todos")
var postTodo = e2e.NewPostAPI[appmodel.TodoForPostRequest, appmodel.TodoForResponse]("localhost:8080", "/api/v1/todos")
var postMail = e2e.NewPostAPI[appmodel.Mail, interface{}]("localhost:8080", "/api/v1/mails")

var getMessages = e2e.NewGetAPI[e2e.Messages]("localhost:8025", "/api/v2/messages")

func Test(t *testing.T) {
	testAppWithoutAuth(t)
	testAuth(t)
	testAppWithAuth(t)
	testMail(t)
}

func testAppWithoutAuth(t *testing.T) {
	// test public is ok
	todos, err := publicGetTodos.Request(false)
	assert.Nil(t, err)
	assert.Equal(t, &[]appmodel.TodoForResponse{}, todos)

	// test private is ng
	todos, err = getTodos.Request(false)
	assert.NotNil(t, err, todos)
}

func testAuth(t *testing.T) {
	// test login
	_, err := login.Request(authmodel.User{Name: "user", Password: "pass"}, false)
	assert.Nil(t, err)

	// test checkLogin
	_, err = checkLogin.Request(false)
	assert.Nil(t, err)
}

func testAppWithAuth(t *testing.T) {
	expected := appmodel.TodoForResponse{ID: 1, Content: "test"}

	// test get todos
	todos, err := getTodos.Request(true)
	assert.Nil(t, err)
	assert.Equal(t, &[]appmodel.TodoForResponse{}, todos)

	// test post todo
	todo, err := postTodo.Request(appmodel.TodoForPostRequest{Content: "test"}, true)
	assert.Nil(t, err)
	assert.Equal(t, &expected, todo)

	// test get todos include new one
	todos, err = getTodos.Request(true)
	assert.Nil(t, err)
	assert.Equal(t, &[]appmodel.TodoForResponse{expected}, todos)
}

func testMail(t *testing.T) {
	// send mail request
	to := "test2@example.com"
	sub := "title"
	msg := "content"
	_, err := postMail.Request(appmodel.Mail{To: to, Sub: &sub, Msg: &msg}, true)
	assert.Nil(t, err)

	// wait to send
	isOK := false
	for i := 0; i < 10; i++ {
		ms, err := getMessages.Request(false)
		assert.Nil(t, err)
		if len(ms.Items) != 1 {
			fmt.Println("waiting email")
			time.Sleep(time.Second)
			continue
		}
		isOK = true
		break
	}
	assert.True(t, isOK)

	// test message
	ms, err := getMessages.Request(false)
	assert.Nil(t, err)
	expectedMs := &e2e.Messages{
		Items: []e2e.Message{
			{
				Raw: e2e.RawMessage{
					From: "test@example.com",
					To:   []string{to},
					Data: fmt.Sprintf("To: %s\r\nSubject: %s\r\n\r\n%s", to, sub, msg),
				},
			},
		},
	}
	assert.Equal(t, expectedMs, ms)
}
