package bls

type ProjectListResult struct {
	Order      *string    `json:"order,omitempty"`
	OrderBy    *string    `json:"orderBy,omitempty"`
	PageNo     *int32     `json:"pageNo,omitempty"`
	PageSize   *int32     `json:"pageSize,omitempty"`
	Total      *int32     `json:"total,omitempty"`
	BlsDefault *Project   `json:"default,omitempty"`
	Projects   []*Project `json:"projects,omitempty"`
}
