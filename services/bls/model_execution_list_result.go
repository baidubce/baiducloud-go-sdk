package bls

type ExecutionListResult struct {
	PageNo     *int32       `json:"pageNo,omitempty"`
	PageSize   *int32       `json:"pageSize,omitempty"`
	TotalCount *int32       `json:"totalCount,omitempty"`
	Executions []*Execution `json:"executions,omitempty"`
}
