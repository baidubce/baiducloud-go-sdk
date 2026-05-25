package vpc

type IpSet struct {
	IpSetId           *string                  `json:"ipSetId,omitempty"`
	Name              *string                  `json:"name,omitempty"`
	Description       *string                  `json:"description,omitempty"`
	IpVersion         *string                  `json:"ipVersion,omitempty"`
	IpAddressInfo     []*TemplateIpAddressInfo `json:"ipAddressInfo,omitempty"`
	BindedInstanceNum *int32                   `json:"bindedInstanceNum,omitempty"`
}
