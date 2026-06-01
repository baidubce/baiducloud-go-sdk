package apm

type DescribeExceptionsRequest struct {
	BeginDatetime *string           `json:"beginDatetime,omitempty"`
	EndDatetime   *string           `json:"endDatetime,omitempty"`
	Service       *string           `json:"service,omitempty"`
	Exceptions    []*ExceptionQuery `json:"exceptions,omitempty"`
}
