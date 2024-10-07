package prometheus_controllers

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/montinger-com/montinger-server/app/shared/models/response_model"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	cpuUsage = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "cpu_usage",
			Help: "CPU usage percentage",
		},
		[]string{"server_name"},
	)
)

func init() {
	prometheus.MustRegister(cpuUsage)
}

func metricsHandler() gin.HandlerFunc {

	randomNumber := rand.Intn(101)
	percentage := float64(randomNumber)
	cpuUsage.WithLabelValues("server-01").Set(percentage)

	fmt.Println("Setting CPU usage to", percentage)

	h := promhttp.Handler()

	return func(c *gin.Context) {
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
		Address: "http://prometheus-production:9090", // Replace with your Prometheus server address
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
		log.Printf("Prometheus query warnings: %v\n", warnings)
	}

	// Return the results as JSON
	c.JSON(http.StatusOK, response_model.Result{Data: result})
}
