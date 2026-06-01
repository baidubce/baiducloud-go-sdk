package apm

type TopoConfig struct {
	Global                  *bool    `json:"global,omitempty"`
	RequestSecondsThreshold *float32 `json:"requestSecondsThreshold,omitempty"`
	ErrorRateThreshold      *float32 `json:"errorRateThreshold,omitempty"`
}
