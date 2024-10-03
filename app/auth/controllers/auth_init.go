package auth_controllers

import (
	"github.com/gin-gonic/gin"
	auth_model "github.com/montinger-com/montinger-server/app/auth/models"
	"github.com/montinger-com/montinger-server/app/utils/validators"
)

func Init(router *gin.Engine) {
	authRoutes := router.Group("/auth")

	authRoutes.POST("/login", validators.ValidateJsonBody[auth_model.LoginDTO], login)
}
