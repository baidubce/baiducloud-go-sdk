package apm

type ServiceTopoConfig struct {
	Global                  *bool    `json:"global,omitempty"`
	RequestSecondsThreshold *float64 `json:"requestSecondsThreshold,omitempty"`
	ErrorRateThreshold      *float64 `json:"errorRateThreshold,omitempty"`
}
