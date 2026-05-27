package csn

type Csn struct {
	CsnId       *string     `json:"csnId,omitempty"`
	Name        *string     `json:"name,omitempty"`
	Description *string     `json:"description,omitempty"`
	Status      *string     `json:"status,omitempty"`
	InstanceNum *int32      `json:"instanceNum,omitempty"`
	CsnBpNum    *int32      `json:"csnBpNum,omitempty"`
	CreateTime  *string     `json:"createTime,omitempty"`
	Tags        []*TagModel `json:"tags,omitempty"`
}
