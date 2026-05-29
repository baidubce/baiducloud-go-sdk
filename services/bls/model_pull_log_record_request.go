package bls

type PullLogRecordRequest struct {
	LogStoreName  *string `json:"-"`
	Project       *string `json:"-"`
	LogStreamName *string `json:"-"`
	StartDateTime *string `json:"-"`
	EndDateTime   *string `json:"-"`
	Limit         *int32  `json:"-"`
	Marker        *string `json:"-"`
}
