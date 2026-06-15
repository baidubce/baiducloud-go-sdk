package bcc

type DeleteAutoRenewRuleRequest struct {
	InstanceId *string `json:"instanceId,omitempty"`
	RenewEip   *bool   `json:"renewEip,omitempty"`
}
