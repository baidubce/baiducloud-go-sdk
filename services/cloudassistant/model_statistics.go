package cloudassistant

type Statistics struct {
	TotalTargetCount  *int32 `json:"totalTargetCount,omitempty"`
	FailedTargetCount *int32 `json:"failedTargetCount,omitempty"`
}
