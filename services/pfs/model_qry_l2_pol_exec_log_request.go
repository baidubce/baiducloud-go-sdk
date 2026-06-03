package pfs

type QryL2PolExecLogRequest struct {
	Action     *string `json:"-"`
	InstanceId *string `json:"instanceId,omitempty"`
	PolicyId   *string `json:"policyId,omitempty"`
	StartTime  *int32  `json:"startTime,omitempty"`
	EndTime    *int32  `json:"endTime,omitempty"`
}
