package finance

type CreateRenewResourceRuleRequest struct {
	AccountId     *string `json:"accountId,omitempty"`
	ServiceType   *string `json:"serviceType,omitempty"`
	Region        *string `json:"region,omitempty"`
	InstanceId    *string `json:"instanceId,omitempty"`
	RenewTimeUnit *string `json:"renewTimeUnit,omitempty"`
	RenewTime     *string `json:"renewTime,omitempty"`
}
