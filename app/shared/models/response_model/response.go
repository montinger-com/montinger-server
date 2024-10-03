package response_model

type Result struct {
	Status      int         `json:"status,omitempty"`
	Environment string      `json:"environment,omitempty"`
	RequestID   interface{} `json:"request_id,omitempty"`
	Path        interface{} `json:"path,omitempty"`
	Timestamp   string      `json:"timestamp,omitempty"`
	Message     string      `json:"message,omitempty"`
	Data        interface{} `json:"data,omitempty"`
	Errors      interface{} `json:"errors,omitempty"`
	Pagination  interface{} `json:"pagination,omitempty"`
}
