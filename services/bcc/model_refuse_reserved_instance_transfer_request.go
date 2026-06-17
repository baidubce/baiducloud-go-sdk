package bcc

type RefuseReservedInstanceTransferRequest struct {
	TransferRecordIds []*string `json:"transferRecordIds,omitempty"`
}
