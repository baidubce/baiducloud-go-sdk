package as

type HealthCheckConfig struct {
	HealthCheckInterval *int32 `json:"healthCheckInterval,omitempty"`
	GraceTime           *int32 `json:"graceTime,omitempty"`
}
