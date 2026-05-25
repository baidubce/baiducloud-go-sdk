package vpc

type CreateIpAddressGroupRequest struct {
	ClientToken   *string                  `json:"-"`
	Name          *string                  `json:"name,omitempty"`
	IpVersion     *string                  `json:"ipVersion,omitempty"`
	IpAddressInfo []*TemplateIpAddressInfo `json:"ipAddressInfo,omitempty"`
	Description   *string                  `json:"description,omitempty"`
}
