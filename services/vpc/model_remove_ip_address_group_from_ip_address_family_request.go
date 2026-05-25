package vpc

type RemoveIpAddressGroupFromIpAddressFamilyRequest struct {
	IpGroupId   *string   `json:"-"`
	ClientToken *string   `json:"-"`
	IpSetIds    []*string `json:"ipSetIds,omitempty"`
}
