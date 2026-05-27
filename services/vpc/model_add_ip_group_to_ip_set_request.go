package vpc

type AddIpGroupToIpSetRequest struct {
	IpGroupId   *string   `json:"-"`
	ClientToken *string   `json:"-"`
	IpSetIds    []*string `json:"ipSetIds,omitempty"`
}
