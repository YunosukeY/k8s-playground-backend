package controller

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	mock_usecase "github.com/YunosukeY/kind-backend/internal/app/mock/usecase"
	"github.com/YunosukeY/kind-backend/internal/app/model"
	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
	zerolog.SetGlobalLevel(zerolog.Disabled)

	m.Run()
}

func TestGetTodos(t *testing.T) {
	cases := []struct {
		name         string
		mockRet      []model.TodoForResponse
		mockErr      error
		expectedCode int
		expectedBody string
	}{
		{
			"0件",
			[]model.TodoForResponse{},
			nil,
			http.StatusOK,
			"[]",
		},
		{
			"1件",
			[]model.TodoForResponse{{ID: 1, Content: "Todo1"}},
			nil,
			http.StatusOK,
			"[{\"id\":1,\"content\":\"Todo1\"}]",
		},
		{
			"異常系",
			nil,
			errors.New(""),
			http.StatusInternalServerError,
			"",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mock_usecase.NewMockUsecase(ctrl)
			mock.EXPECT().GetAllTodos(gomock.Any()).Return(tt.mockRet, tt.mockErr)
			con := NewController(util.NewTestTracer(), mock)

			res := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(res)
			req, _ := http.NewRequest(http.MethodGet, "/api/v1/todos", nil)
			ctx.Request = req
			con.GetTodos(ctx)

			assert.Equal(t, tt.expectedCode, res.Code)
			assert.Equal(t, tt.expectedBody, res.Body.String())
		})
	}
}

func TestPostTodo(t *testing.T) {
	cases := []struct {
		name         string
		input        interface{}
		mockRet      *model.TodoForResponse
		mockError    error
		expectedCode int
		expectedBody string
	}{
		{
			"正常系：正常",
			model.TodoForPostRequest{Content: "todo1"},
			&model.TodoForResponse{ID: 1, Content: "todo1"},
			nil,
			http.StatusOK,
			"{\"id\":1,\"content\":\"todo1\"}",
		},
		{
			"正常系：余計なフィールド",
			model.TodoForResponse{ID: 10, Content: "todo1"},
			&model.TodoForResponse{ID: 1, Content: "todo1"},
			nil,
			http.StatusOK,
			"{\"id\":1,\"content\":\"todo1\"}",
		},
		{
			"異常系：不正なContent",
			struct{}{},
			nil,
			nil,
			http.StatusBadRequest,
			"",
		},
		{
			"異常系：DBエラー",
			model.TodoForPostRequest{Content: "todo"},
			nil,
			errors.New(""),
			http.StatusInternalServerError,
			"",
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mock_usecase.NewMockUsecase(ctrl)
			if tt.expectedCode != http.StatusBadRequest {
				mock.EXPECT().CreateTodo(gomock.Any(), gomock.Any()).Return(tt.mockRet, tt.mockError)
			}
			con := NewController(util.NewTestTracer(), mock)

			res := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(res)
			body, _ := json.Marshal(tt.input)
			ctx.Request, _ = http.NewRequest(
				http.MethodPost,
				"/api/v1/todos",
				bytes.NewReader(body),
			)
			con.PostTodo(ctx)

			assert.Equal(t, tt.expectedCode, res.Code)
			assert.Equal(t, tt.expectedBody, res.Body.String())
		})
	}
}

func TestPostMail(t *testing.T) {
	cases := []struct {
		name         string
		input        interface{}
		mockError    error
		expectedCode int
	}{
		{
			"正常系：正常",
			model.Mail{To: "test@example.com", Sub: util.ToPtr(""), Msg: util.ToPtr("")},
			nil,
			http.StatusOK,
		},
		{
			"異常系：メールアドレスでない",
			model.Mail{To: "example.com", Sub: util.ToPtr(""), Msg: util.ToPtr("")},
			nil,
			http.StatusBadRequest,
		},
		{
			"異常系：SubとMsgがない",
			model.Mail{To: "test@example.com"},
			nil,
			http.StatusBadRequest,
		},
		{
			"異常系：キューエラー",
			model.Mail{To: "test@example.com", Sub: util.ToPtr(""), Msg: util.ToPtr("")},
			errors.New(""),
			http.StatusInternalServerError,
		},
	}

	for _, tt := range cases {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			mock := mock_usecase.NewMockUsecase(ctrl)
			if tt.expectedCode != http.StatusBadRequest {
				mock.EXPECT().SendMail(gomock.Any(), gomock.Any()).Return(tt.mockError)
			}
			con := NewController(util.NewTestTracer(), mock)

			res := httptest.NewRecorder()
			ctx, _ := gin.CreateTestContext(res)
			body, _ := json.Marshal(tt.input)
			ctx.Request, _ = http.NewRequest(
				http.MethodPost,
				"/api/v1/mails",
				bytes.NewReader(body),
			)
			con.PostMail(ctx)

			assert.Equal(t, tt.expectedCode, res.Code)
		})
	}
}
