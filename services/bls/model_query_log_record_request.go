package bls

type QueryLogRecordRequest struct {
	LogStoreName  *string `json:"-"`
	Project       *string `json:"-"`
	LogStreamName *string `json:"-"`
	Query         *string `json:"-"`
	StartDateTime *string `json:"-"`
	EndDateTime   *string `json:"-"`
	Marker        *string `json:"-"`
	Limit         *int32  `json:"-"`
	Sort          *string `json:"-"`
	PageNo        *int32  `json:"-"`
	PageSize      *int32  `json:"-"`
}
