package bcc

type ModifyReservedInstanceOrder struct {
	ReservedInstanceId *string `json:"reservedInstanceId,omitempty"`
	OrderId            *string `json:"orderId,omitempty"`
}
