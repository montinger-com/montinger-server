package monitors_model

import "time"

type Monitor struct {
	ID         string     `json:"id,omitempty" bson:"_id,omitempty"`
	Name       string     `json:"name" bson:"name"`
	Type       string     `json:"type" bson:"type"`
	Status     string     `json:"status" bson:"status"`
	LastDataOn *time.Time `json:"last_data_on,omitempty" bson:"last_data_on,omitempty"`
	LastData   LastData   `json:"last_data,omitempty" bson:"last_data,omitempty"`
	APIKey     string     `json:"api_key,omitempty" bson:"api_key,omitempty"`
}

type LastData struct {
	CPUUsage    float64 `json:"cpu_usage,omitempty" bson:"cpu_usage,omitempty"`
	MemoryUsage float64 `json:"memory_usage,omitempty" bson:"memory_usage,omitempty"`
}

type MonitorCreateDTO struct {
	Name string `json:"name" validate:"required"`
	Type string `json:"type" validate:"required"`
}

type MonitorCreateResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Status string `json:"status"`
	Token  string `json:"token"`
}

type MonitorRegisterDTO struct {
	ID    string `json:"id" validate:"required"`
	Token string `json:"token" validate:"required"`
}

type MonitorRegisterResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Type   string `json:"type"`
	Status string `json:"status"`
	APIKey string `json:"api_key"`
}

type MonitorPushDTO struct {
	LastData LastData `json:"last_data,omitempty" bson:"last_data,omitempty"`
}

type IDParamDTO struct {
	ID string `json:"monitor_id" validate:"required" uri:"monitor_id" binding:"required"`
}
