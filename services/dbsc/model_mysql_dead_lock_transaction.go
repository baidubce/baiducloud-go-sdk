package dbsc

type MysqlDeadLockTransaction struct {
	Database       *string `json:"database,omitempty"`
	HeapNo         *int32  `json:"heapNo,omitempty"`
	Index          *string `json:"index,omitempty"`
	IsPrediction   *bool   `json:"isPrediction,omitempty"`
	LockMode       *string `json:"lockMode,omitempty"`
	LockType       *string `json:"lockType,omitempty"`
	PageNo         *int32  `json:"pageNo,omitempty"`
	RecordLockType *string `json:"recordLockType,omitempty"`
	SpaceId        *int32  `json:"spaceId,omitempty"`
	Table          *string `json:"table,omitempty"`
	WaitHold       *string `json:"waitHold,omitempty"`
	Query          *string `json:"query,omitempty"`
	ThreadID       *int32  `json:"threadID,omitempty"`
	TrxId          *string `json:"trxId,omitempty"`
	TrxSeq         *int32  `json:"trxSeq,omitempty"`
	TrxTime        *string `json:"trxTime,omitempty"`
	User           *string `json:"user,omitempty"`
	Victim         *int32  `json:"victim,omitempty"`
}
