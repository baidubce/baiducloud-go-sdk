package apm

type ServiceRequestConfig struct {
	Global                            *bool    `json:"global,omitempty"`
	ServerSlowRequestThresholdSeconds *float64 `json:"serverSlowRequestThresholdSeconds,omitempty"`
	DbSlowRequestThresholdSeconds     *float64 `json:"dbSlowRequestThresholdSeconds,omitempty"`
	OkHttpStatus                      []*int32 `json:"okHttpStatus,omitempty"`
}
