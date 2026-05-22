package blb

type CreateAppBlbRequest struct {
	ClientToken                  *string           `json:"-"`
	Name                         *string           `json:"name,omitempty"`
	BlbType                      *string           `json:"type,omitempty"`
	Desc                         *string           `json:"desc,omitempty"`
	SubnetId                     *string           `json:"subnetId,omitempty"`
	VpcId                        *string           `json:"vpcId,omitempty"`
	Address                      *string           `json:"address,omitempty"`
	Eip                          *string           `json:"eip,omitempty"`
	Tags                         []*TagModel       `json:"tags,omitempty"`
	Billing                      *BillingForCreate `json:"billing,omitempty"`
	PerformanceLevel             *string           `json:"performanceLevel,omitempty"`
	AutoRenewLength              *int32            `json:"autoRenewLength,omitempty"`
	AutoRenewTimeUnit            *string           `json:"autoRenewTimeUnit,omitempty"`
	ResourceGroupId              *string           `json:"resourceGroupId,omitempty"`
	AllowDelete                  *bool             `json:"allowDelete,omitempty"`
	AllowModify                  *bool             `json:"allowModify,omitempty"`
	ModificationProtectionReason *string           `json:"modificationProtectionReason,omitempty"`
	AllocateIpv6                 *bool             `json:"allocateIpv6,omitempty"`
}
