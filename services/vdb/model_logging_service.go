package vdb

type LoggingService struct {
	Enabled   *bool   `json:"enabled,omitempty"`
	LogType   *string `json:"logType,omitempty"`
	Status    *string `json:"status,omitempty"`
	Supported *bool   `json:"supported,omitempty"`
}
