package bls

type ListDownloadTaskRequest struct {
	Project      *string `json:"project,omitempty"`
	LogStoreName *string `json:"logStoreName,omitempty"`
	OrderBy      *string `json:"orderBy,omitempty"`
	Order        *string `json:"order,omitempty"`
	PageNo       *int32  `json:"pageNo,omitempty"`
	PageSize     *int32  `json:"pageSize,omitempty"`
}
