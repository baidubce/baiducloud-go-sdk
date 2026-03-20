package vpc

type Cgw struct {
	CgwId       *string `json:"cgwId,omitempty"`
	Name        *string `json:"name,omitempty"`
	Ip          *string `json:"ip,omitempty"`
	Description *string `json:"description,omitempty"`
	CreatedTime *string `json:"createdTime,omitempty"`
	UpdatedTime *string `json:"updatedTime,omitempty"`
}
