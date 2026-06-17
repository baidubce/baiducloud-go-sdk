package bcc

type ModifyReservedInstancesRequest struct {
	ReservedInstances []*ReservedInstance `json:"reservedInstances,omitempty"`
}
