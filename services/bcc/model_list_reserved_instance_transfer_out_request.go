package bcc

type ListReservedInstanceTransferOutRequest struct {
	Marker              *string   `json:"-"`
	MaxKeys             *int32    `json:"-"`
	ReservedInstanceIds []*string `json:"reservedInstanceIds,omitempty"`
	TransferRecordIds   []*string `json:"transferRecordIds,omitempty"`
	Spec                *string   `json:"spec,omitempty"`
	Status              *string   `json:"status,omitempty"`
}
