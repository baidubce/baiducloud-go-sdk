package bcc

type CreateReservedInstancesRequest struct {
	ReservedInstanceName     *string     `json:"reservedInstanceName,omitempty"`
	Scope                    *string     `json:"scope,omitempty"`
	ZoneName                 *string     `json:"zoneName,omitempty"`
	Spec                     *string     `json:"spec,omitempty"`
	OfferingType             *string     `json:"offeringType,omitempty"`
	InstanceCount            *int32      `json:"instanceCount,omitempty"`
	ReservedInstanceCount    *int32      `json:"reservedInstanceCount,omitempty"`
	ReservedInstanceTime     *int32      `json:"reservedInstanceTime,omitempty"`
	ReservedInstanceTimeUnit *string     `json:"reservedInstanceTimeUnit,omitempty"`
	AutoRenew                *bool       `json:"autoRenew,omitempty"`
	AutoRenewTimeUnit        *string     `json:"autoRenewTimeUnit,omitempty"`
	AutoRenewTime            *int32      `json:"autoRenewTime,omitempty"`
	EffectiveTime            *string     `json:"effectiveTime,omitempty"`
	Tags                     []*TagModel `json:"tags,omitempty"`
	TicketId                 *string     `json:"ticketId,omitempty"`
	EhcClusterId             *string     `json:"ehcClusterId,omitempty"`
}
