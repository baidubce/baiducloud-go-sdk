package scs

type GetHotGroupListRequest struct {
	PageSize *int32 `json:"pageSize,omitempty"`
	PageNo   *int32 `json:"pageNo,omitempty"`
}
