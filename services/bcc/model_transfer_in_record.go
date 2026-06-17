package bcc

type TransferInRecord struct {
	TransferRecordId     *string               `json:"transferRecordId,omitempty"`
	GrantorUserId        *string               `json:"grantorUserId,omitempty"`
	Status               *string               `json:"status,omitempty"`
	ReservedInstanceInfo *ReservedInstanceInfo `json:"reservedInstanceInfo,omitempty"`
	ApplicationTime      *string               `json:"applicationTime,omitempty"`
	ExpireTime           *string               `json:"expireTime,omitempty"`
	EndTime              *string               `json:"endTime,omitempty"`
}
