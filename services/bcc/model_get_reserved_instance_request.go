package bcc

type GetReservedInstanceRequest struct {
	Marker                 *string   `json:"-"`
	MaxKeys                *int32    `json:"-"`
	ReservedInstanceIds    []*string `json:"reservedInstanceIds,omitempty"`
	ReservedInstanceName   *string   `json:"reservedInstanceName,omitempty"`
	ZoneName               *string   `json:"zoneName,omitempty"`
	ReservedInstanceStatus *string   `json:"reservedInstanceStatus,omitempty"`
	Spec                   *string   `json:"spec,omitempty"`
	OfferingType           *string   `json:"offeringType,omitempty"`
	OsType                 *string   `json:"osType,omitempty"`
	InstanceId             *string   `json:"instanceId,omitempty"`
	InstanceName           *string   `json:"instanceName,omitempty"`
	IsDeduct               *bool     `json:"isDeduct,omitempty"`
	EhcClusterId           *string   `json:"ehcClusterId,omitempty"`
	SortKey                *string   `json:"sortKey,omitempty"`
	SortDir                *string   `json:"sortDir,omitempty"`
	ReservedInstanceSource *string   `json:"reservedInstanceSource,omitempty"`
	Scope                  *string   `json:"scope,omitempty"`
}
