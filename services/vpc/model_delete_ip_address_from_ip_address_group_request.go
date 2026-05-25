package vpc

type DeleteIpAddressFromIpAddressGroupRequest struct {
	IpSetId       *string   `json:"-"`
	ClientToken   *string   `json:"-"`
	IpAddressInfo []*string `json:"ipAddressInfo,omitempty"`
}
