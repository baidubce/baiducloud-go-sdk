package bcc

type RevokeReservedInstanceTransferRequest struct {
	TransferRecordIds []*string `json:"transferRecordIds,omitempty"`
}
