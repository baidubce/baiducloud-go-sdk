package vpc

type AddIpAddressToIpGroupRequest struct {
	IpSetId       *string                  `json:"-"`
	ClientToken   *string                  `json:"-"`
	IpAddressInfo []*TemplateIpAddressInfo `json:"ipAddressInfo,omitempty"`
}
