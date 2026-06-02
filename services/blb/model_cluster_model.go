package blb

type ClusterModel struct {
	Id         *string `json:"id,omitempty"`
	Name       *string `json:"name,omitempty"`
	BlbType    *string `json:"type,omitempty"`
	Status     *string `json:"status,omitempty"`
	CcuCount   *int32  `json:"ccuCount,omitempty"`
	CreateTime *string `json:"createTime,omitempty"`
	ExpireTime *string `json:"expireTime,omitempty"`
	Desc       *string `json:"desc,omitempty"`
}
