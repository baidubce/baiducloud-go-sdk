package eip

type CancelEipTransferRequest struct {
	ClientToken    *string   `json:"-"`
	TransferIdList []*string `json:"transferIdList,omitempty"`
}
