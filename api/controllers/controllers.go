package controllers

import (
	"github.com/gin-gonic/gin"
	auth_controllers "github.com/montinger-com/montinger-server/app/auth/controllers"
)

func Init(router *gin.Engine) {
	auth_controllers.Init(router)
}
