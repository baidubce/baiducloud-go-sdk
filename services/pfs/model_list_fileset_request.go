package pfs

type ListFilesetRequest struct {
	Action      *string `json:"-"`
	InstanceId  *string `json:"instanceId,omitempty"`
	FilesetId   *string `json:"filesetId,omitempty"`
	FilesetName *string `json:"filesetName,omitempty"`
	Manner      *string `json:"manner,omitempty"`
	Order       *string `json:"order,omitempty"`
	OrderBy     *string `json:"orderBy,omitempty"`
	PageNo      *int32  `json:"pageNo,omitempty"`
	PageSize    *int32  `json:"pageSize,omitempty"`
}
