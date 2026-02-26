package eip

type ReceiveEipTransferRequest struct {
	ClientToken    *string   `json:"-"`
	TransferIdList []*string `json:"transferIdList,omitempty"`
}
