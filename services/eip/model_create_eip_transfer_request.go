

package eip

type CreateEipTransferRequest struct {
    ClientToken *string `json:"-"`
    TransferType *string `json:"transferType,omitempty"`
    TransferResourceList []*string `json:"transferResourceList,omitempty"`
    ToUserId *string `json:"toUserId,omitempty"`
}