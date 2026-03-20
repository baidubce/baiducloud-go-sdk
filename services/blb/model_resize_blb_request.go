package blb

type ResizeBlbRequest struct {
	BlbId            *string `json:"-"`
	ClientToken      *string `json:"-"`
	PerformanceLevel *string `json:"performanceLevel,omitempty"`
}
