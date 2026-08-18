package finance

type RenewResource struct {
	ServiceType      *string `json:"serviceType,omitempty"`
	Region           *string `json:"region,omitempty"`
	ShortId          *string `json:"shortId,omitempty"`
	AccountId        *string `json:"accountId,omitempty"`
	InstanceId       *string `json:"instanceId,omitempty"`
	ExpireTime       *string `json:"expireTime,omitempty"`
	AloneRenewEnable *bool   `json:"aloneRenewEnable,omitempty"`
	AlreadyRenewSet  *bool   `json:"alreadyRenewSet,omitempty"`
	RenewTimeUnit    *string `json:"renewTimeUnit,omitempty"`
	RenewTime        *string `json:"renewTime,omitempty"`
}
