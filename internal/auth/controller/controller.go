package controller

import (
	"net/http"

	"github.com/YunosukeY/kind-backend/internal/auth/model"
	"github.com/YunosukeY/kind-backend/internal/auth/repository"
	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/otel/trace"
)

const authHeaderValue = "authorized"
const authCookieKey = "username"

type Controller interface {
	GetSession(ctx *gin.Context)
	PostSession(ctx *gin.Context)
}

type controller struct {
	t trace.Tracer
	c repository.Cache
}

func NewController(t trace.Tracer, c repository.Cache) Controller {
	return controller{t, c}
}

func (c controller) GetSession(ctx *gin.Context) {
	child, span := c.t.Start(ctx.Request.Context(), util.FuncName())
	defer span.End()

	session := sessions.Default(ctx)
	username := session.Get(authCookieKey)
	log.Debug().Interface("username", username).Msg("")
	if username == nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	_, err := c.c.Get(child, username.(string))
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.Header(util.AuthHeaderKey, authHeaderValue)
	ctx.Status(http.StatusOK)
}

// TODO: 本当は正しいユーザか確認が必要
// TODO: 本当はユーザ名とパスワードではなくUUIDとユーザ名をセットする必要
func (c controller) PostSession(ctx *gin.Context) {
	child, span := c.t.Start(ctx.Request.Context(), util.FuncName())
	defer span.End()

	var u model.User
	if err := ctx.ShouldBindJSON(&u); err != nil {
		log.Error().Err(err).Msg("")
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	log.Debug().Interface("user", u).Msg("")

	if err := c.c.Set(child, u.Name, u.Password); err != nil {
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	session := sessions.Default(ctx)
	session.Set(authCookieKey, u.Name)
	err := session.Save()
	if err != nil {
		log.Error().Err(err).Msg("")
		c.c.Delete(child, u.Name) // nolint
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.Status(http.StatusOK)
}
