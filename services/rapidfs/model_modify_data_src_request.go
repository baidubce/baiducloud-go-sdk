package rapidfs

type ModifyDataSrcRequest struct {
	ClientToken *string `json:"-"`
	DataSrcId   *string `json:"dataSrcId,omitempty"`
	InstanceId  *string `json:"instanceId,omitempty"`
	KeepSymlink *bool   `json:"keepSymlink,omitempty"`
	AuthGroupId *string `json:"authGroupId,omitempty"`
	Description *string `json:"description,omitempty"`
	QuotaMiB    *int32  `json:"quotaMiB,omitempty"`
}
