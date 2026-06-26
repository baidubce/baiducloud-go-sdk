package bls

type QueryLogHistogramRequest struct {
	LogStoreName  *string `json:"-"`
	Project       *string `json:"-"`
	LogStreamName *string `json:"-"`
	LogStoreType  *string `json:"-"`
	Query         *string `json:"-"`
	StartDateTime *string `json:"-"`
	EndDateTime   *string `json:"-"`
}
