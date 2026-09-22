package scs

type DelayInfoConsoleItem struct {
	SourceCluster *string `json:"sourceCluster,omitempty"`
	DestCluster   *string `json:"destCluster,omitempty"`
	DelayResult   *int64  `json:"delayResult,omitempty"`
	TimeResult    *int64  `json:"timeResult,omitempty"`
}
