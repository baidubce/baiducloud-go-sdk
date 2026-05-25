package vpc

type AddIpAddressesToTheIpAddressGroupRequest struct {
	IpSetId       *string                  `json:"-"`
	ClientToken   *string                  `json:"-"`
	IpAddressInfo []*TemplateIpAddressInfo `json:"ipAddressInfo,omitempty"`
}
