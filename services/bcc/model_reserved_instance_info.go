package bcc

type ReservedInstanceInfo struct {
	ReservedInstanceId     *string     `json:"reservedInstanceId,omitempty"`
	ReservedInstanceUuid   *string     `json:"reservedInstanceUuid,omitempty"`
	ReservedInstanceName   *string     `json:"reservedInstanceName,omitempty"`
	Scope                  *string     `json:"scope,omitempty"`
	ZoneName               *string     `json:"zoneName,omitempty"`
	LogicalZone            *string     `json:"logicalZone,omitempty"`
	Spec                   *string     `json:"spec,omitempty"`
	ReservedType           *string     `json:"reservedType,omitempty"`
	OfferingType           *string     `json:"offeringType,omitempty"`
	OsType                 *string     `json:"osType,omitempty"`
	ReservedInstanceStatus *string     `json:"reservedInstanceStatus,omitempty"`
	InstanceCount          *int32      `json:"instanceCount,omitempty"`
	EffectiveTime          *string     `json:"effectiveTime,omitempty"`
	ExpireTime             *string     `json:"expireTime,omitempty"`
	TransferInTime         *string     `json:"transferInTime,omitempty"`
	AutoRenew              *bool       `json:"autoRenew,omitempty"`
	RenewTimeUnit          *string     `json:"renewTimeUnit,omitempty"`
	RenewTime              *int32      `json:"renewTime,omitempty"`
	NextRenewTime          *string     `json:"nextRenewTime,omitempty"`
	FlavorSubType          *string     `json:"flavorSubType,omitempty"`
	ProductCategory        *string     `json:"productCategory,omitempty"`
	IsNeedEhcCluster       *bool       `json:"isNeedEhcCluster,omitempty"`
	Tags                   []*TagModel `json:"tags,omitempty"`
	InstanceId             *string     `json:"instanceId,omitempty"`
	InstanceIds            []*string   `json:"instanceIds,omitempty"`
	InstanceName           *string     `json:"instanceName,omitempty"`
	InstanceNames          []*string   `json:"instanceNames,omitempty"`
	TransferIn             *bool       `json:"transferIn,omitempty"`
	GrantorUserId          *string     `json:"grantorUserId,omitempty"`
	TimeGranularity        *string     `json:"timeGranularity,omitempty"`
	ReservedInstanceTime   *int32      `json:"reservedInstanceTime,omitempty"`
	EhcClusterId           *string     `json:"ehcClusterId,omitempty"`
}
