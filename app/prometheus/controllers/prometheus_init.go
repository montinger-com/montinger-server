package prometheus_controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func Init(router *gin.Engine) {
	router.POST("/metrics", metricsHandler)
	router.GET("/query", queryPrometheus)
	router.GET("/metrics/prometheus", gin.WrapH(promhttp.Handler()))
}
