package bcc

type RenewReservedInstanceRequest struct {
	ReservedInstanceIds      []*string `json:"reservedInstanceIds,omitempty"`
	ReservedInstanceTime     *string   `json:"reservedInstanceTime,omitempty"`
	ReservedInstanceTimeUnit *string   `json:"reservedInstanceTimeUnit,omitempty"`
	AutoRenew                *bool     `json:"autoRenew,omitempty"`
	AutoRenewTimeUnit        *string   `json:"autoRenewTimeUnit,omitempty"`
	AutoRenewTime            *int32    `json:"autoRenewTime,omitempty"`
}
