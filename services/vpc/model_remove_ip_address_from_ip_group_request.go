package vpc

type RemoveIpAddressFromIpGroupRequest struct {
	IpSetId       *string   `json:"-"`
	ClientToken   *string   `json:"-"`
	IpAddressInfo []*string `json:"ipAddressInfo,omitempty"`
}
