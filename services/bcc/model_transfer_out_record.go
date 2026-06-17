package bcc

type TransferOutRecord struct {
	TransferRecordId     *string               `json:"transferRecordId,omitempty"`
	RecipientUserId      *string               `json:"recipientUserId,omitempty"`
	Status               *string               `json:"status,omitempty"`
	ReservedInstanceInfo *ReservedInstanceInfo `json:"reservedInstanceInfo,omitempty"`
	ApplicationTime      *string               `json:"applicationTime,omitempty"`
	ExpireTime           *string               `json:"expireTime,omitempty"`
	EndTime              *string               `json:"endTime,omitempty"`
}
