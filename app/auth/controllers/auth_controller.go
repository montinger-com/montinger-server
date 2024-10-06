package auth_controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	auth_model "github.com/montinger-com/montinger-server/app/auth/models"
	auth_services "github.com/montinger-com/montinger-server/app/auth/services"
	"github.com/montinger-com/montinger-server/app/shared/enums"
	"github.com/montinger-com/montinger-server/app/shared/models/response_model"
	users_model "github.com/montinger-com/montinger-server/app/users/models"
	users_services "github.com/montinger-com/montinger-server/app/users/services"
	"github.com/montinger-com/montinger-server/lib/exceptions"

	"github.com/montinger-com/montinger-server/app/utils/helpers"
)

var authService auth_services.AuthService
var usersService users_services.UsersService

func init() {
	authService = *auth_services.NewAuthService()
	usersService = *users_services.NewUserService()
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

func register(c *gin.Context) {
	registerDTO := helpers.GetJsonBody[auth_model.RegisterDTO](c)

	if helpers.IsEmpty(registerDTO.Email) || helpers.IsEmpty(registerDTO.Password) {
		c.JSON(http.StatusBadRequest, response_model.Result{Message: exceptions.InvalidInput.Error(), Errors: []string{exceptions.InvalidInput.Error()}})
		return
	}

	existingUser, _ := usersService.GetByEmail(registerDTO.Email)

	if existingUser != nil && existingUser.Status != enums.Deleted {
		c.JSON(http.StatusBadRequest, response_model.Result{Message: exceptions.UserAlreadyExists.Error(), Errors: []string{exceptions.UserAlreadyExists.Error()}})
		return
	}

	user, err := usersService.Create(&users_model.User{Email: registerDTO.Email, Password: registerDTO.Password})
	if err != nil {
		c.JSON(http.StatusInternalServerError, response_model.Result{Message: err.Error(), Errors: []string{err.Error()}})
		return
	}

	c.JSON(http.StatusCreated, response_model.Result{Data: user})
}
