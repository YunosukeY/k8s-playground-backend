package controller

import (
	"net/http"

	"github.com/YunosukeY/kind-backend/internal/app/model"
	"github.com/YunosukeY/kind-backend/internal/app/usecase"
	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
)

type Controller interface {
	GetTodos(ctx *gin.Context)
	PostTodo(ctx *gin.Context)
	PostMail(ctx *gin.Context)
}

type controller struct {
	t trace.Tracer
	u usecase.Usecase
}

func NewController(t trace.Tracer, u usecase.Usecase) Controller {
	return controller{t, u}
}

func (c controller) GetTodos(ctx *gin.Context) {
	child, span := c.t.Start(ctx.Request.Context(), util.FuncName())
	defer span.End()

	todos, err := c.u.GetAllTodos(child)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	log.Debug().Interface("todos", todos).Msg("")
	ctx.JSON(http.StatusOK, todos)
}

func (c controller) PostTodo(ctx *gin.Context) {
	child, span := c.t.Start(ctx.Request.Context(), util.FuncName())
	defer span.End()

	var todo model.TodoForPostRequest
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		log.Error().Err(err).Msg("")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	todoWithID, err := c.u.CreateTodo(child, todo)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	log.Debug().Interface("todo", todoWithID).Msg("")

	ctx.JSON(http.StatusOK, todoWithID)
}

func (c controller) PostMail(ctx *gin.Context) {
	child, span := c.t.Start(ctx.Request.Context(), util.FuncName())
	defer span.End()

	var mail model.Mail
	if err := ctx.ShouldBindJSON(&mail); err != nil {
		log.Error().Err(err).Msg("")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	log.Debug().Interface("mail", mail).Msg("")

	if err := c.u.SendMail(child, mail); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}
