package vpc

type RemoveIpGroupFromIpSetRequest struct {
	IpGroupId   *string   `json:"-"`
	ClientToken *string   `json:"-"`
	IpSetIds    []*string `json:"ipSetIds,omitempty"`
}
