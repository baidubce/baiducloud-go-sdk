package eip

type RejectEipTransferRequest struct {
	ClientToken    *string   `json:"-"`
	TransferIdList []*string `json:"transferIdList,omitempty"`
}
