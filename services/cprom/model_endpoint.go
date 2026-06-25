package cprom

type Endpoint struct {
	Port               *string `json:"port,omitempty"`
	Path               *string `json:"path,omitempty"`
	Interval           *string `json:"interval,omitempty"`
	MatchedTargetCount *int32  `json:"matchedTargetCount,omitempty"`
}
