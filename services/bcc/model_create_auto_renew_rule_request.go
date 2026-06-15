package bcc

type CreateAutoRenewRuleRequest struct {
	InstanceId    *string `json:"instanceId,omitempty"`
	RenewTimeUnit *string `json:"renewTimeUnit,omitempty"`
	RenewTime     *int32  `json:"renewTime,omitempty"`
	RenewEip      *bool   `json:"renewEip,omitempty"`
}
