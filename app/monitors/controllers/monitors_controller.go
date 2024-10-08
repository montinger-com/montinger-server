package monitors_controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	monitors_model "github.com/montinger-com/montinger-server/app/monitors/models"
	monitors_services "github.com/montinger-com/montinger-server/app/monitors/services"
	"github.com/montinger-com/montinger-server/app/shared/models/response_model"
	"github.com/montinger-com/montinger-server/app/utils/helpers"
	"github.com/montinger-com/montinger-server/lib/exceptions"
)

var monitorsService monitors_services.MonitorsService

func init() {
	monitorsService = *monitors_services.NewMonitorsService()
}

func create(c *gin.Context) {
	monitorDTO := helpers.GetJsonBody[monitors_model.MonitorCreateDTO](c)

	tokens, err := monitorsService.Create(&monitors_model.Monitor{Name: monitorDTO.Name, Type: monitorDTO.Type})

	if err != nil {
		c.JSON(http.StatusBadRequest, response_model.Result{Message: exceptions.InvalidUsernameOrPassword.Error(), Errors: []string{err.Error()}})
		return
	}

	c.JSON(http.StatusOK, response_model.Result{Data: tokens})
}

func register(c *gin.Context) {
	monitorDTO := helpers.GetJsonBody[monitors_model.MonitorRegisterDTO](c)

	monitor, err := monitorsService.Register(&monitorDTO)
	if err != nil {
		c.JSON(http.StatusBadRequest, response_model.Result{Message: exceptions.RegistrationFailed.Error(), Errors: []string{err.Error()}})
		return
	}

	c.JSON(http.StatusOK, response_model.Result{Data: monitor})
}

func push(c *gin.Context) {
	params := helpers.GetParams[monitors_model.IDParamDTO](c)
	monitorDTO := helpers.GetJsonBody[monitors_model.MonitorPushDTO](c)
	apiKey := helpers.GetAPIKey(c)

	err := monitorsService.Push(params.ID, &monitorDTO, apiKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, response_model.Result{Message: exceptions.FailedToPushDeviceData.Error(), Errors: []string{err.Error()}})
		return
	}

	c.JSON(http.StatusOK, response_model.Result{Message: "data pushed successfully"})
}
