package vpc

type CreateIpSetRequest struct {
	ClientToken *string   `json:"-"`
	Name        *string   `json:"name,omitempty"`
	IpVersion   *string   `json:"ipVersion,omitempty"`
	IpSetIds    []*string `json:"ipSetIds,omitempty"`
	Description *string   `json:"description,omitempty"`
}
