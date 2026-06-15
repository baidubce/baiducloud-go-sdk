package bcc

type ChangeToPrepaidRequest struct {
	InstanceId      *string `json:"-"`
	Duration        *int32  `json:"duration,omitempty"`
	AutoRenew       *bool   `json:"autoRenew,omitempty"`
	AutoRenewPeriod *int32  `json:"autoRenewPeriod,omitempty"`
	RelationCds     *bool   `json:"relationCds,omitempty"`
}
