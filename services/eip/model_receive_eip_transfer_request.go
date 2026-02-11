

package eip

type ReceiveEipTransferRequest struct {
    Action *string `json:"-"`
    ClientToken *string `json:"-"`
    TransferIdList []*string `json:"transferIdList,omitempty"`
}