package as

type CreateAsGroupV2Request struct {
	GroupName         *string            `json:"groupName,omitempty"`
	ZoneInfo          []*ZoneInfo        `json:"zoneInfo,omitempty"`
	Config            *GroupConfig       `json:"config,omitempty"`
	KeypairId         *string            `json:"keypairId,omitempty"`
	KeypairName       *string            `json:"keypairName,omitempty"`
	KeepImageLogin    *int32             `json:"keepImageLogin,omitempty"`
	Blb               []*BlbInfo         `json:"blb,omitempty"`
	BlbUnbindWaitTime *int32             `json:"blbUnbindWaitTime,omitempty"`
	Nodes             []*NodeInfo        `json:"nodes,omitempty"`
	Eip               *EipInfo           `json:"eip,omitempty"`
	EipConfig         *EipConfig         `json:"eipConfig,omitempty"`
	Billing           *BillingInfo       `json:"billing,omitempty"`
	Rds               []*string          `json:"rds,omitempty"`
	Scs               []*string          `json:"scs,omitempty"`
	HealthCheck       *HealthCheckConfig `json:"healthCheck,omitempty"`
	ExpansionStrategy *string            `json:"expansionStrategy,omitempty"`
	ShrinkageStrategy *string            `json:"shrinkageStrategy,omitempty"`
	AssignTagInfo     *AssignTagInfo     `json:"assignTagInfo,omitempty"`
	CmdConfig         *CmdConfig         `json:"cmdConfig,omitempty"`
	BccNameConfig     *BccNameConfig     `json:"bccNameConfig,omitempty"`
}
