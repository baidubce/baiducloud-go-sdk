package bls

type ListLogStoreViewRequest struct {
	Project  *string `json:"project,omitempty"`
	Name     *string `json:"name,omitempty"`
	Order    *string `json:"order,omitempty"`
	OrderBy  *string `json:"orderBy,omitempty"`
	PageNo   *int32  `json:"pageNo,omitempty"`
	PageSize *int32  `json:"pageSize,omitempty"`
}
