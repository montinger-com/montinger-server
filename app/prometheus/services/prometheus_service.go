package prometheus_services

import (
	"context"
	"fmt"
	"time"

	"github.com/montinger-com/montinger-server/config"
	"github.com/prometheus/client_golang/api"
	v1 "github.com/prometheus/client_golang/api/prometheus/v1"
	"github.com/prometheus/common/model"
	"github.com/rashintha/logger"
)

type PrometheusService struct {
}

func NewPrometheusService() *PrometheusService {

	return &PrometheusService{}
}

func (s *PrometheusService) GetDataByMetric(metric string, timePeriod int) (model.Value, error) {

	// Create a new API client for Prometheus
	client, err := api.NewClient(api.Config{
		Address: fmt.Sprintf("http://%v:%v", config.PROMETHEUS_HOST, config.PROMETHEUS_PORT),
	})
	if err != nil {
		logger.Errorf("Error creating Prometheus API client: %v\n", err)
		return nil, err
	}

	v1api := v1.NewAPI(client)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	query := fmt.Sprintf("%v[%vm]", metric, timePeriod)

	// Query Prometheus using the v1 API
	result, warnings, err := v1api.Query(ctx, query, time.Now())
	if err != nil {
		logger.Errorf("Error querying Prometheus: %v\n", err)
		return nil, err
	}

	if len(warnings) > 0 {
		logger.Warningf("Prometheus query warnings: %v\n", warnings)
	}

	return result, nil
}
