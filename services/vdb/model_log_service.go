package vdb

type LogService struct {
	Enabled       *bool   `json:"enabled,omitempty"`
	LogStoreName  *string `json:"logStoreName,omitempty"`
	Project       *string `json:"project,omitempty"`
	RetentionDays *int32  `json:"retentionDays,omitempty"`
}
