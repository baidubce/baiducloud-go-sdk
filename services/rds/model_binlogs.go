package rds

type Binlogs struct {
	BinlogId          *string `json:"binlogId,omitempty"`
	BinlogSizeInBytes *int32  `json:"binlogSizeInBytes,omitempty"`
	BinlogEndTime     *string `json:"binlogEndTime,omitempty"`
	BinlogStartTime   *string `json:"binlogStartTime,omitempty"`
	BinlogStatus      *string `json:"binlogStatus,omitempty"`
	BinlogRealId      *string `json:"binlogRealId,omitempty"`
}
