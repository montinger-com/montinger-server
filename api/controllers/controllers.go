package controllers

import (
	"github.com/gin-gonic/gin"
	auth_controllers "github.com/montinger-com/montinger-server/app/auth/controllers"
	monitors_controllers "github.com/montinger-com/montinger-server/app/monitors/controllers"
	prometheus_controllers "github.com/montinger-com/montinger-server/app/prometheus/controllers"
)

func Init(router *gin.Engine) {
	auth_controllers.Init(router)
	monitors_controllers.Init(router)
	prometheus_controllers.Init(router)
}
