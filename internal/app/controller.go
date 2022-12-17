package app

import (
	"net/http"

	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
)

type Controller interface {
	getTodos(ctx *gin.Context)
	postTodo(ctx *gin.Context)
	postMail(ctx *gin.Context)
}

type controller struct {
	t trace.Tracer
	r Repository
	q Queue
}

func NewController(t trace.Tracer, r Repository, q Queue) Controller {
	return controller{t, r, q}
}

func (c controller) getTodos(ctx *gin.Context) {
	child, span := c.t.Start(ctx.Request.Context(), util.FuncName())
	defer span.End()

	todos, err := c.r.findAllTodos(child)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	log.Debug().Interface("todos", todos).Msg("")
	ctx.JSON(http.StatusOK, todos)
}

func (c controller) postTodo(ctx *gin.Context) {
	child, span := c.t.Start(ctx.Request.Context(), util.FuncName())
	defer span.End()

	var todo TodoForPostRequest
	if err := ctx.ShouldBindJSON(&todo); err != nil {
		log.Error().Err(err).Msg("")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}

	todoWithID, err := c.r.createTodo(child, todo)
	if err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	log.Debug().Interface("todo", todoWithID).Msg("")

	ctx.JSON(http.StatusOK, todoWithID)
}

func (c controller) postMail(ctx *gin.Context) {
	child, span := c.t.Start(ctx.Request.Context(), util.FuncName())
	defer span.End()

	var mail Mail
	if err := ctx.ShouldBindJSON(&mail); err != nil {
		log.Error().Err(err).Msg("")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	log.Debug().Interface("mail", mail).Msg("")

	if err := c.q.push(child, mail); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}
