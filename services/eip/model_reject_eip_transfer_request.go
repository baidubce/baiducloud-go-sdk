

package eip

type RejectEipTransferRequest struct {
    Action *string `json:"-"`
    ClientToken *string `json:"-"`
    TransferIdList []*string `json:"transferIdList,omitempty"`
}