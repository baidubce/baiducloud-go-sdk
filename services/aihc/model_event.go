package aihc

type Event struct {
	Reason         *string `json:"reason,omitempty"`
	Message        *string `json:"message,omitempty"`
	FirstTimestamp *string `json:"firstTimestamp,omitempty"`
	LastTimestamp  *string `json:"lastTimestamp,omitempty"`
	Count          *int32  `json:"count,omitempty"`
	AihcType       *int32  `json:"type,omitempty"`
}
