package bls

type PolicyListResult struct {
	PageNo     *int32    `json:"pageNo,omitempty"`
	PageSize   *int32    `json:"pageSize,omitempty"`
	TotalCount *int32    `json:"totalCount,omitempty"`
	Policies   []*Policy `json:"policies,omitempty"`
}
