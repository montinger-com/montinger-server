package prometheus_controllers

import (
	"github.com/gin-gonic/gin"
)

func Init(router *gin.Engine) {
	router.GET("/query", queryPrometheus)
	router.GET("/metrics", metricsHandler())
}
