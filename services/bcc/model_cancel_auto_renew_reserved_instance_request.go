package bcc

type CancelAutoRenewReservedInstanceRequest struct {
	ReservedInstanceIds []*string `json:"reservedInstanceIds,omitempty"`
}
