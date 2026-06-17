package bcc

type AutoRenewReservedInstanceRequest struct {
	ReservedInstanceIds []*string `json:"reservedInstanceIds,omitempty"`
	AutoRenewTimeUnit   *string   `json:"autoRenewTimeUnit,omitempty"`
	AutoRenewTime       *int32    `json:"autoRenewTime,omitempty"`
}
