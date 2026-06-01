package apm

type RequestConfig struct {
	Global                            *bool    `json:"global,omitempty"`
	ServerSlowRequestThresholdSeconds *float32 `json:"serverSlowRequestThresholdSeconds,omitempty"`
	DbSlowRequestThresholdSeconds     *float32 `json:"dbSlowRequestThresholdSeconds,omitempty"`
	OkHttpStatus                      []*int32 `json:"okHttpStatus,omitempty"`
}
