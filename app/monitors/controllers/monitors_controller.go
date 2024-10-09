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

func getAll(c *gin.Context) {
	monitors, err := monitorsService.GetAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, response_model.Result{Message: exceptions.FailedFetchingData.Error(), Errors: []string{err.Error()}})
		return
	}

	monitorsResponse := make([]*monitors_model.MonitorResponse, 0)
	for _, monitor := range monitors {
		monitorsResponse = append(monitorsResponse, &monitors_model.MonitorResponse{
			ID:         monitor.ID,
			Name:       monitor.Name,
			Type:       monitor.Type,
			Status:     monitor.Status,
			LastDataOn: monitor.LastDataOn,
			LastData:   monitor.LastData,

			CreatedAt: monitor.CreatedAt,
			UpdatedAt: monitor.UpdatedAt,
		})
	}

	c.JSON(http.StatusOK, response_model.Result{Data: monitorsResponse})
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
		c.JSON(http.StatusBadRequest, response_model.Result{Message: exceptions.FailedToPushData.Error(), Errors: []string{err.Error()}})
		return
	}

	c.JSON(http.StatusOK, response_model.Result{Message: "data pushed successfully"})
}

func getData(c *gin.Context) {
	query := helpers.GetJsonQuery[monitors_model.MonitorDataQueryParamDTO](c)

	monitors, err := monitorsService.GetDataByMetrics(query.Types, query.TimePeriod)
	if err != nil {
		c.JSON(http.StatusBadRequest, response_model.Result{Message: exceptions.FailedFetchingData.Error(), Errors: []string{err.Error()}})
		return
	}

	c.JSON(http.StatusOK, response_model.Result{Data: monitors})
}
