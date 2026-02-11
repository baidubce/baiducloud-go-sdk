

package eip

type CancelEipTransferRequest struct {
    Action *string `json:"-"`
    ClientToken *string `json:"-"`
    TransferIdList []*string `json:"transferIdList,omitempty"`
}