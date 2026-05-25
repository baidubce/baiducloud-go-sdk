package vpc

type AddIpAddressGroupToIpAddressFamilyRequest struct {
	IpGroupId   *string   `json:"-"`
	ClientToken *string   `json:"-"`
	IpSetIds    []*string `json:"ipSetIds,omitempty"`
}
