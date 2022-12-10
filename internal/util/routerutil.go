package util

import (
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/logger"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

const AuthHeaderKey = "X-Auth"

func loggerMiddleware() gin.HandlerFunc {
	customLogger := func(c *gin.Context, l zerolog.Logger) zerolog.Logger {
		return zerolog.New(os.Stdout).
			With().
			Timestamp().
			Int("status", c.Writer.Status()).
			Str("method", c.Request.Method).
			Str("path", c.Request.URL.Path).
			Str("ip", c.ClientIP()).
			Str("user_agent", c.Request.UserAgent()).
			Str(AuthHeaderKey, c.Request.Header.Get(AuthHeaderKey)).
			Logger()
	}
	return logger.SetLogger(
		logger.WithLogger(customLogger),
		logger.WithSkipPath([]string{"/healthz", "/metrics"}),
	)
}

func DefaultRouter() *gin.Engine {
	router := gin.New()
	router.Use(cors.Default(), gin.Recovery(), loggerMiddleware())

	return router
}
