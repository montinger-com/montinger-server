package monitors_controllers

import (
	"github.com/gin-gonic/gin"
	monitors_model "github.com/montinger-com/montinger-server/app/monitors/models"
	"github.com/montinger-com/montinger-server/app/utils/validators"
)

func Init(router *gin.Engine) {
	monitorsRoutes := router.Group("/monitors")

	monitorsRoutes.POST("", validators.ValidateJsonBody[monitors_model.MonitorCreateDTO], create)
	monitorsRoutes.GET("", getAll)
	monitorsRoutes.POST("/register", validators.ValidateJsonBody[monitors_model.MonitorRegisterDTO], register)
	monitorsRoutes.POST("/:monitor_id/push", validators.ValidateJsonBody[monitors_model.MonitorPushDTO], validators.ValidatePathParams[monitors_model.IDParamDTO], push)
	monitorsRoutes.GET("/data", validators.ValidateQueryParams[monitors_model.MonitorDataQueryParamDTO], getData)
}
