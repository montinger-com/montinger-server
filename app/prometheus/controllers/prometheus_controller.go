package prometheus_controllers

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	monitors_service "github.com/montinger-com/montinger-server/app/monitors/services"
	"github.com/montinger-com/montinger-server/app/shared/models/response_model"
	"github.com/montinger-com/montinger-server/config"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rashintha/logger"
)

var monitorsService monitors_service.MonitorsService

var (
	cpuUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "cpu_usage",
			Help: "CPU usage percentage",
		}, []string{"server_name"},
	)

	memoryUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "memory_usage",
			Help: "Memory usage percentage",
		}, []string{"server_name"},
	)
)

func init() {
	prometheus.MustRegister(cpuUsage)
	prometheus.MustRegister(memoryUsage)

	monitorsService = *monitors_service.NewMonitorsService()
}

func metricsHandler() gin.HandlerFunc {
	h := promhttp.Handler()

	return func(c *gin.Context) {

		monitors, err := monitorsService.GetAll()

		if err != nil {
			logger.Errorf("Error getting monitors: %v\n", err)
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "Error getting monitors"})
			return
		}

		for _, monitor := range monitors {
			if monitor.Type == "server" && monitor.LastData != nil && monitor.LastDataOn != nil && !monitor.LastDataOn.Add(1*time.Minute).After(time.Now()) {
				memoryUsage.WithLabelValues(monitor.ID).Set(monitor.LastData.MemoryUsage)
				cpuUsage.WithLabelValues(monitor.ID).Set(monitor.LastData.CPUUsage)
			}
		}

		h.ServeHTTP(c.Writer, c.Request)
	}
}

func queryPrometheus(c *gin.Context) {
	// Get the query parameter from the request
	query := c.Query("query")
	if query == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing 'query' parameter"})
		return
	}

	// Create a new API client for Prometheus
	client, err := api.NewClient(api.Config{
		Address: fmt.Sprintf("http://%v:%v", config.PROMETHEUS_HOST, config.PROMETHEUS_PORT),
	})
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("error creating Prometheus client: %w", err))
		return
	}

	v1api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Query Prometheus using the v1 API
	result, warnings, err := v1api.Query(ctx, query, time.Now())
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, fmt.Errorf("error querying Prometheus: %w", err))
		return
	}

	if len(warnings) > 0 {
		logger.Warningf("Prometheus query warnings: %v\n", warnings)
	}

	// Return the results as JSON
	c.JSON(http.StatusOK, response_model.Result{Data: result})
}
