package bls

type ListLogStoreRequest struct {
	Project     *string `json:"project,omitempty"`
	NamePattern *string `json:"namePattern,omitempty"`
	Order       *string `json:"order,omitempty"`
	OrderBy     *string `json:"orderBy,omitempty"`
	PageNo      *int32  `json:"pageNo,omitempty"`
	PageSize    *int32  `json:"pageSize,omitempty"`
}
