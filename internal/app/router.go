package app

import (
	"net/http"

	"github.com/YunosukeY/kind-backend/internal/app/controller"
	"github.com/YunosukeY/kind-backend/internal/util"
	"github.com/gin-gonic/gin"
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

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := c.Request.Header.Get(util.AuthHeaderKey)
		if auth == "" {
			c.Status(http.StatusUnauthorized)
			c.Abort()
		}
		c.Next()
	}
}

func (r router) handler() *gin.Engine {
	router := util.DefaultRouter()
	router.Use(otelgin.Middleware(""))

	v1 := router.Group("/api/v1")
	{
		private := v1.Group("/")
		{
			private.Use(authMiddleware())
			private.GET("todos", r.c.GetTodos)
			private.POST("todos", r.c.PostTodo)
			private.POST("mails", r.c.PostMail)
		}

		public := v1.Group("/public")
		{
			public.GET("todos", r.c.GetTodos)
		}
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
		r, shutdownProvider = initializeRouter("app")
	} else {
		r, shutdownProvider = initializeDummyRouter("app")
	}
	defer shutdownProvider()

	go util.RunPodCommonHandler()
	if err := r.handler().Run(); err != nil {
		log.Panic().Err(err).Msg("")
		panic(err)
	}
}
