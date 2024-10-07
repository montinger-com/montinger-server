package monitors_model

import "time"

type Monitor struct {
	ID         string     `json:"id" bson:"_id,omitempty"`
	Name       string     `json:"name" bson:"name"`
	Type       string     `json:"type" bson:"type"`
	Status     string     `json:"status" bson:"status"`
	LastDataOn *time.Time `json:"last_data_on" bson:"last_data_on"`
	LastData   LastData   `json:"last_data" bson:"last_data"`
}

type LastData struct {
	CPUUsage    float64 `json:"cpu_usage" bson:"cpu_usage"`
	MemoryUsage float64 `json:"memory_usage" bson:"memory_usage"`
}
