package auth

import (
	"github.com/YunosukeY/kind-backend/internal/auth/controller"
	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/securecookie"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

type router struct {
	c controller.Controller
}

func newRouter(c controller.Controller) router {
	return router{c}
}

func sessionMiddleware() gin.HandlerFunc {
	store := cookie.NewStore(securecookie.GenerateRandomKey(32), securecookie.GenerateRandomKey(16))
	store.Options(sessions.Options{MaxAge: 60 * 60 * 24, HttpOnly: true})
	return sessions.Sessions("session", store)
}

func (r router) handler() *gin.Engine {
	router := util.DefaultRouter()
	router.Use(sessionMiddleware(), otelgin.Middleware(""))

	v1 := router.Group("/api/v1")
	{
		v1.GET("sessions", r.c.GetSession)
		v1.POST("sessions", r.c.PostSession)

	}

	return router
}

func Run(dummy bool) {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
	log.Logger = log.With().Caller().Logger()
	gin.SetMode(gin.ReleaseMode)

	var r router
	var shutdownProvider func()
	if !dummy {
		r, shutdownProvider = initializeRouter("auth")
	} else {
		r, shutdownProvider = initializeDummyRouter("auth")
	}
	defer shutdownProvider()

	go util.RunPodCommonHandler()
	if err := r.handler().Run(); err != nil {
		log.Panic().Err(err).Msg("")
		panic(err)
	}
}
