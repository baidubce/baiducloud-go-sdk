package bcc

type ElasticityAssuranceModel struct {
	ElasticityAssuranceId   *string          `json:"elasticityAssuranceId,omitempty"`
	ElasticityAssuranceName *string          `json:"ElasticityAssuranceName,omitempty"`
	Spec                    *string          `json:"spec,omitempty"`
	Status                  *string          `json:"status,omitempty"`
	ZoneName                *string          `json:"zoneName,omitempty"`
	Capacity                *int32           `json:"capacity,omitempty"`
	UsedCapacity            *int32           `json:"usedCapacity,omitempty"`
	AvailableCapacity       *int32           `json:"availableCapacity,omitempty"`
	CreateTime              *string          `json:"createTime,omitempty"`
	ExpireTime              *string          `json:"expireTime,omitempty"`
	AutoRenew               *bool            `json:"autoRenew,omitempty"`
	AutoRenewPeriod         *int32           `json:"autoRenewPeriod,omitempty"`
	Tags                    []*Tag           `json:"tags,omitempty"`
	Instances               []*InstanceModel `json:"instances,omitempty"`
}
