package bls

type Target struct {
	Query                 *string   `json:"query,omitempty"`
	StartTimeOffsetMinute *int32    `json:"startTimeOffsetMinute,omitempty"`
	EndTimeOffsetMinute   *int32    `json:"endTimeOffsetMinute,omitempty"`
	Object                *LogStore `json:"object,omitempty"`
}
