package bls

type ListProjectRequest struct {
	Name        *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	OrderBy     *string `json:"orderBy,omitempty"`
	Order       *string `json:"order,omitempty"`
	PageNo      *int32  `json:"pageNo,omitempty"`
	PageSize    *int32  `json:"pageSize,omitempty"`
}
