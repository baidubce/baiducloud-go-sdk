package apm

type DescribeDbStatementRequest struct {
	BeginDatetime *string           `json:"beginDatetime,omitempty"`
	EndDatetime   *string           `json:"endDatetime,omitempty"`
	Service       *string           `json:"service,omitempty"`
	Statements    []*StatementQuery `json:"statements,omitempty"`
}
