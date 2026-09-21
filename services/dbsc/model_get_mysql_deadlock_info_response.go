package dbsc

import "github.com/baidubce/baiducloud-go-sdk/bce"

type GetMysqlDeadlockInfoResponse struct {
	bce.BaseResponse
	MetaInfo         *interface{}                `json:"metaInfo,omitempty"`
	DeadLockId       *string                     `json:"deadLockId,omitempty"`
	Timestamp        *string                     `json:"timestamp,omitempty"`
	RawContent       *string                     `json:"rawContent,omitempty"`
	TransactionLocks []*MysqlDeadLockTransaction `json:"transactionLocks,omitempty"`
}
