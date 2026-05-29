package bls

type Statistics struct {
	ExecutionTimeInMs *int32 `json:"executionTimeInMs,omitempty"`
	ScanCount         *int32 `json:"scanCount,omitempty"`
	ScanBytes         *int32 `json:"scanBytes,omitempty"`
}
