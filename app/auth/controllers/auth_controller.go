package auth_controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	auth_model "github.com/montinger-com/montinger-server/app/auth/models"
	auth_services "github.com/montinger-com/montinger-server/app/auth/services"
	"github.com/montinger-com/montinger-server/app/shared/models/response_model"
	"github.com/montinger-com/montinger-server/lib/exceptions"

	"github.com/montinger-com/montinger-server/app/utils/helpers"
)

var authService auth_services.AuthService

func init() {
	authService = *auth_services.NewAuthService()
}

func login(c *gin.Context) {
	loginDTO := helpers.GetJsonBody[auth_model.LoginDTO](c)
	token, err := authService.AuthenticateUser(loginDTO.Email, loginDTO.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, response_model.Result{Message: exceptions.InvalidUsernameOrPassword.Error(), Errors: []string{err.Error()}})
		return
	}

	c.JSON(http.StatusOK, response_model.Result{Data: token})
}
