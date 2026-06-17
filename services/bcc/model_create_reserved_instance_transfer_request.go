package bcc

type CreateReservedInstanceTransferRequest struct {
	ReservedInstanceIds []*string `json:"reservedInstanceIds,omitempty"`
	RecipientAccountId  *string   `json:"recipientAccountId,omitempty"`
}
