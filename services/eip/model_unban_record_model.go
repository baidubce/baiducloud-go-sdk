package eip

type UnbanRecordModel struct {
	Ip            *string `json:"ip,omitempty"`
	ProtectType   *int32  `json:"protectType,omitempty"`
	StartTime     *string `json:"startTime,omitempty"`
	ExpectEndTime *string `json:"expectEndTime,omitempty"`
	Status        *int32  `json:"status,omitempty"`
}
