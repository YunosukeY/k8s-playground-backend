package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func healthCheck(ctx *gin.Context) {
	ctx.Status(http.StatusOK)
}

func podCommonHandler() *gin.Engine {
	router := DefaultRouter()

	router.GET("/healthz", healthCheck)
	router.GET("/metrics", gin.WrapH(promhttp.Handler()))

	return router
}

func RunPodCommonHandler() {
	r := podCommonHandler()
	if err := r.Run(":8888"); err != nil {
		panic(err)
	}
}
