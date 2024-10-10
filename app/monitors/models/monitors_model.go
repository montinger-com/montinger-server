package monitors_model

import "time"

type Monitor struct {
	ID     string `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string `json:"name" bson:"name"`
	Type   string `json:"type" bson:"type"`
	Status string `json:"status" bson:"status"`

	APIKey string `json:"api_key,omitempty" bson:"api_key,omitempty"`

	LastDataOn *time.Time `json:"last_data_on,omitempty" bson:"last_data_on,omitempty"`
	LastData   *LastData  `json:"last_data,omitempty" bson:"last_data,omitempty"`

	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" bson:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" bson:"deleted_at"`
}

type MonitorResponse struct {
	ID     string `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string `json:"name" bson:"name"`
	Type   string `json:"type" bson:"type"`
	Status string `json:"status" bson:"status"`

	LastDataOn *time.Time `json:"last_data_on,omitempty" bson:"last_data_on,omitempty"`
	LastData   *LastData  `json:"last_data,omitempty" bson:"last_data,omitempty"`

	CreatedAt time.Time  `json:"created_at" bson:"created_at"`
	UpdatedAt *time.Time `json:"updated_at,omitempty" bson:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at,omitempty" bson:"deleted_at"`
}

type LastData struct {
	CPU    *CPU    `json:"cpu,omitempty" bson:"cpu,omitempty"`
	Memory *Memory `json:"memory,omitempty" bson:"memory,omitempty"`
	OS     *OS     `json:"os,omitempty" bson:"os,omitempty"`
	Uptime uint64  `json:"uptime,omitempty" bson:"uptime,omitempty"`
}

type CPU struct {
	UsedPercent float64 `json:"used_percent,omitempty" bson:"used_percent,omitempty"`
}

type Memory struct {
	Total       uint64  `json:"total,omitempty" bson:"total,omitempty"`
	Available   uint64  `json:"available,omitempty" bson:"available,omitempty"`
	Used        uint64  `json:"used,omitempty" bson:"used,omitempty"`
	UsedPercent float64 `json:"used_percent,omitempty" bson:"used_percent,omitempty"`
}

type OS struct {
	Type            string `json:"type,omitempty" bson:"type,omitempty"`
	Platform        string `json:"platform,omitempty" bson:"platform,omitempty"`
	PlatformFamily  string `json:"platform_family,omitempty" bson:"platform_family,omitempty"`
	PlatformVersion string `json:"platform_version,omitempty" bson:"platform_version,omitempty"`
	KernelVersion   string `json:"kernel_version,omitempty" bson:"kernel_version,omitempty"`
	KernelArch      string `json:"kernel_arch,omitempty" bson:"kernel_arch,omitempty"`
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

type MonitorDataQueryParamDTO struct {
	Types      []string `json:"types" form:"types" binding:"required"`
	TimePeriod int      `json:"time_period" form:"time_period" binding:"required"`
}

type MonitorDataResponse struct {
	ID          string       `json:"id"`
	TimePeriod  *TimePeriod  `json:"time_period,omitempty"`
	MemoryUsage *MemoryUsage `json:"memory_usage,omitempty"`
	CPUUsage    *CPUUsage    `json:"cpu_usage,omitempty"`
}

type TimePeriod struct {
	Duration *int    `json:"duration,omitempty"`
	Unit     *string `json:"unit,omitempty"`
}

type MemoryUsage struct {
	Unit   *string         `json:"unit,omitempty"`
	Values []UsageIntValue `json:"values,omitempty"`
}

type UsageIntValue struct {
	Time  *time.Time `json:"time,omitempty"`
	Value float64    `json:"value,omitempty"`
}

type CPUUsage struct {
	Unit   *string         `json:"unit,omitempty"`
	Values []UsageIntValue `json:"values,omitempty"`
}
