package bls

type DownloadTaskListResult struct {
	Tasks    []*Task `json:"tasks,omitempty"`
	Total    *int32  `json:"total,omitempty"`
	OrderBy  *string `json:"orderBy,omitempty"`
	Order    *string `json:"order,omitempty"`
	PageNo   *int32  `json:"pageNo,omitempty"`
	PageSize *int32  `json:"pageSize,omitempty"`
}
