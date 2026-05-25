package vpc

type IpGroup struct {
	IpGroupId         *string   `json:"ipGroupId,omitempty"`
	Name              *string   `json:"name,omitempty"`
	Description       *string   `json:"description,omitempty"`
	IpVersion         *string   `json:"ipVersion,omitempty"`
	IpSetIds          []*string `json:"ipSetIds,omitempty"`
	BindedInstanceNum *int32    `json:"bindedInstanceNum,omitempty"`
}
