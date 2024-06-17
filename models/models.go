package models

type Response struct {
	Fields     map[string]string `json:"fields,omitempty"`
	Status     string            `json:"status,omitempty"`
	Success    bool              `json:"success,omitempty"`
	Message    string            `json:"message,omitempty"`
	StatusCode int               `json:"status_code,omitempty"`
	Data       interface{}       `json:"data,omitempty"`
}