package bcc

type AcceptReservedInstanceTransferRequest struct {
	TransferRecordId *string `json:"transferRecordId,omitempty"`
	EhcClusterId     *string `json:"ehcClusterId,omitempty"`
}
