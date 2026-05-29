package bls

type PushLogRecordRequest struct {
	LogStoreName  *string      `json:"-"`
	Project       *string      `json:"-"`
	LogStreamName *string      `json:"logStreamName,omitempty"`
	BlsType       *string      `json:"type,omitempty"`
	LogRecords    []*LogRecord `json:"logRecords,omitempty"`
	Tags          []*LogTag    `json:"tags,omitempty"`
}
