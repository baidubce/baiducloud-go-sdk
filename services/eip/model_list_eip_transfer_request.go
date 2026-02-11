

package eip

type ListEipTransferRequest struct {
    MaxKeys *string `json:"-"`
    Marker *string `json:"-"`
    Direction *string `json:"-"`
    TransferId *string `json:"-"`
    Status *string `json:"-"`
    FuzzyTransferId *string `json:"-"`
    FuzzyInstanceId *string `json:"-"`
    FuzzyInstanceName *string `json:"-"`
    FuzzyInstanceIp *string `json:"-"`
}