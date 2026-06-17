package bls

type ExecutionSummaryResult struct {
	ExecutionStats []*ExecutionStats `json:"executionStats,omitempty"`
	PageNo         *int32            `json:"pageNo,omitempty"`
	PageSize       *int32            `json:"pageSize,omitempty"`
	TotalCount     *int32            `json:"totalCount,omitempty"`
}
